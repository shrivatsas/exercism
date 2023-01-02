package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
	"sync"
)

type coord struct {
	x int
	y int
	z int
}

type Particle struct {
	id int
	p  coord
	v  coord
	a  coord
	d  int
}

func (c *coord) add(oc coord) {
	c.x += oc.x
	c.y += oc.y
	c.z += oc.z
}

func (c coord) dist() int {
	return c.x + c.y + c.z
}

func (p *Particle) move() {
	p.v.add(p.a)
	p.p.add(p.v)
	p.d = p.p.dist()
}

func move(particle Particle, iterations int, pch chan Particle, wg *sync.WaitGroup) {
	for i := 0; i < iterations; i++ {
		particle.move()
	}

	pch <- particle
	wg.Done()
}

func closestBuf(particles []Particle) Particle {
	var cp Particle
	var wg sync.WaitGroup
	wg.Add(len(particles))
	pch := make(chan Particle, len(particles))

	var findcp = func(p Particle) {
		if cp.d == 0 || p.d < cp.d {
			cp = p
		}
	}
	for _, particle := range particles {
		go move(particle, 10000, pch, &wg)
	}

	wg.Wait()
	close(pch)

	for p := range pch {
		findcp(p)
	}

	return cp
}

func closestUnbuf(particles []Particle) Particle {
	var cp Particle
	var wg sync.WaitGroup
	wg.Add(len(particles))
	pch := make(chan Particle)
	defer close(pch)

	var findcp = func(p Particle) {
		if cp.d == 0 || p.d < cp.d {
			cp = p
		}
	}
	go func() {
		for {
			select {
			case p := <-pch:
				findcp(p)
			}
		}
	}()

	for _, particle := range particles {
		go move(particle, 10000, pch, &wg)
	}

	wg.Wait()
	return cp
}

func main() {
	file, _ := os.Open("particles.txt")
	scanner := bufio.NewScanner(file)

	var particles []Particle

	var i int
	for scanner.Scan() {
		p, v, a := parse(scanner.Text())
		pr := newParticle(i, p, v, a)
		particles = append(particles, *pr)
		i++
	}

	c := closestBuf(particles)
	log.Println("[Buffered]Closest particle is: ", c.id)

	d := closestUnbuf(particles)
	log.Println("[Unbuffered]Closest particle is: ", d.id)
}

func parse(pva string) (coord, coord, coord) {
	parts := strings.Fields(pva)
	p := parsecoord(parts[0])
	v := parsecoord(parts[1])
	a := parsecoord(parts[2])
	return p, v, a
}

func parsecoord(s string) coord {
	p := strings.Split(s[3:len(s)-1], ",")
	x, _ := strconv.Atoi(p[0])
	y, _ := strconv.Atoi(p[1])
	z, _ := strconv.Atoi(p[2])
	c := coord{
		x: x,
		y: y,
		z: z,
	}
	return c
}

func newParticle(i int, p coord, v coord, a coord) *Particle {
	return &Particle{
		id: i,
		p:  p,
		v:  v,
		a:  a,
	}
}
