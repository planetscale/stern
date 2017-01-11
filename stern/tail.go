//   Copyright 2016 Wercker Holding BV
//
//   Licensed under the Apache License, Version 2.0 (the "License");
//   you may not use this file except in compliance with the License.
//   You may obtain a copy of the License at
//
//       http://www.apache.org/licenses/LICENSE-2.0
//
//   Unless required by applicable law or agreed to in writing, software
//   distributed under the License is distributed on an "AS IS" BASIS,
//   WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//   See the License for the specific language governing permissions and
//   limitations under the License.

package stern

import (
	"bufio"
	"context"
	"crypto/md5"
	"fmt"
	"regexp"

	"github.com/fatih/color"
	"github.com/pkg/errors"

	corev1 "k8s.io/client-go/1.5/kubernetes/typed/core/v1"
	"k8s.io/client-go/1.5/pkg/api/v1"
	"k8s.io/client-go/1.5/rest"
)

type Tail struct {
	PodName        string
	ContainerName  string
	Options        *TailOptions
	req            *rest.Request
	closed         chan struct{}
	podColor       *color.Color
	containerColor *color.Color
}

type TailOptions struct {
	Timestamps   bool
	SinceSeconds int64
	Exclude      []*regexp.Regexp
}

// NewTail returns a new tail for a Kubernetes container inside a pod
func NewTail(podName, containerName string, options *TailOptions) *Tail {
	return &Tail{
		PodName:       podName,
		ContainerName: containerName,
		Options:       options,
		closed:        make(chan struct{}),
	}
}

// var index = 0

var colorList = [][2]*color.Color{
	{color.New(color.FgHiCyan), color.New(color.FgCyan)},
	{color.New(color.FgHiGreen), color.New(color.FgGreen)},
	{color.New(color.FgHiMagenta), color.New(color.FgMagenta)},
	{color.New(color.FgHiYellow), color.New(color.FgYellow)},
	{color.New(color.FgHiBlue), color.New(color.FgBlue)},
	{color.New(color.FgHiRed), color.New(color.FgRed)},
}

func Color256() []*color.Color {
	fgbg := color.Attribute(38)
	someConst := color.Attribute(5)
	out := []*color.Color{}
	for green := 0; green < 6; green++ {
		for red := 0; red < 6; red++ {
			for blue := 0; blue < 6; blue++ {
				myCol := color.Attribute(16 + (red * 36) + (green * 6) + blue)
				out = append(out, color.New(fgbg, someConst, myCol))
			}
		}
	}
	return out
}

// Start starts tailing
func (t *Tail) Start(ctx context.Context, i corev1.PodInterface) {

	sum := md5.Sum([]byte(t.PodName))
	index := int(sum[0])

	cols := Color256()

	colorIndex := index % len(cols)
	t.podColor = cols[colorIndex]
	t.containerColor = cols[colorIndex]

	go func() {
		g := color.New(color.FgHiGreen, color.Bold).SprintFunc()
		p := t.podColor.SprintFunc()
		c := t.podColor.SprintFunc()
		fmt.Printf("%s %s › %s\n", g("+"), p(t.PodName), c(t.ContainerName))

		req := i.GetLogs(t.PodName, &v1.PodLogOptions{
			Follow:       true,
			Timestamps:   t.Options.Timestamps,
			Container:    t.ContainerName,
			SinceSeconds: &t.Options.SinceSeconds,
		})

		stream, err := req.Stream()
		if err != nil {
			fmt.Println(errors.Wrapf(err, "Error opening stream to %s: %s\n", t.PodName, t.ContainerName))
			return
		}
		defer stream.Close()

		go func() {
			<-t.closed
			stream.Close()
		}()

		reader := bufio.NewReader(stream)

	OUTER:
		for {
			line, err := reader.ReadBytes('\n')
			if err != nil {
				return
			}

			str := string(line)

			for _, rex := range t.Options.Exclude {
				if rex.MatchString(str) {
					continue OUTER
				}
			}

			t.Print(str)
		}
	}()

	go func() {
		<-ctx.Done()
		close(t.closed)
	}()
}

// Close stops tailing
func (t *Tail) Close() {
	r := color.New(color.FgHiRed, color.Bold).SprintFunc()
	p := t.podColor.SprintFunc()
	fmt.Printf("%s %s\n", r("-"), p(t.PodName))
	close(t.closed)
}

// Print prints a color coded log message with the pod and container names
func (t *Tail) Print(msg string) {
	p := t.podColor.SprintFunc()
	c := t.podColor.SprintFunc()
	fmt.Printf("%s %s %s", p(t.PodName), c(t.ContainerName), msg)
}
