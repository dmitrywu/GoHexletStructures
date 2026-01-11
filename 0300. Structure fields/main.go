package main

import "fmt"

type Post struct {
	ID     int
	Title  string
	Author string
	Likes  int
}

type Product struct {
	Name  string
	Price float64
	Stock int
	Tags  []string
}

type Cart struct {
	Owner   string
	Items   []Product
	Coupons map[string]int
}

type Candidate struct {
	Name            string
	ExperienceYears int
	Skills          []string
	Remote          bool
	DesiredSalary   int
}

func HasSkill(c Candidate, skill string) bool {
	for _, v := range c.Skills {
		// fmt.Println(i, v, c.Skills[i])
		if v == skill {
			return true
		}
	}
	return false
}

func DescribeCandidate(c Candidate) string {
	remoteCheck := "yes"
	if c.Remote == false {
		remoteCheck = "no"
	}
	return fmt.Sprintf("%s: %dy exp, salary: %d, remote: %s", c.Name, c.ExperienceYears, c.DesiredSalary, remoteCheck)
}

func main() {
	fmt.Print("\033[H\033[2J")

	posts := []Post{
		{ID: 1, Title: "Go vs Python", Author: "Ivan", Likes: 120},
		{ID: 2, Title: "GC. It's worked?", Author: "Maria", Likes: 45},
	}

	fmt.Println(posts[1].Author)

	c := Candidate{
		Name:            "Irina",
		ExperienceYears: 5,
		Skills:          []string{"Go", "Docker", "Kubernetes"},
		Remote:          true,
		DesiredSalary:   4200,
	}

	fmt.Println(DescribeCandidate(c))
	// Irina: 5y exp, salary: 4200, remote: yes
	fmt.Println(HasSkill(c, "Go"))
	// true
	fmt.Println(HasSkill(c, "Rust"))
	// false

}
