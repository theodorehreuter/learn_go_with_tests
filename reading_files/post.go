package blogposts

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"strings"
)

type Post struct {
	Title       string
	Description string
	Body        string
	Tags        []string
}

const (
	titleSeparator       = "Title: "
	descriptionSeparator = "Description: "
	tagSeparator         = "Tags: "
)

func newPost(postFile io.Reader) (Post, error) {
	scanner := bufio.NewScanner(postFile)

	readMetaLine := func(tagName string) string {
		scanner.Scan()
		return strings.TrimPrefix(scanner.Text(), tagName)
	}
	title := readMetaLine(titleSeparator)
	description := readMetaLine(descriptionSeparator)
	tags := strings.Split(readMetaLine(tagSeparator), ", ")
	body := getBody(scanner)

	return Post{
		Title:       title,
		Description: description,
		Tags:        tags,
		Body:        body,
	}, nil
}

func getBody(scanner *bufio.Scanner) string {

	scanner.Scan()

	buf := bytes.Buffer{}
	for scanner.Scan() {
		fmt.Fprintln(&buf, scanner.Text())
	}

	body := strings.TrimSuffix(buf.String(), "\n")

	return body
}
