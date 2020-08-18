package main

import (
	"strconv"
)

//Defining the Footer object
type Footer struct {
	List []string `json:"list"`
}

//function responsible to create the footer pagination
func newFooter(currentpage, around, totalpages, boundaries int) Footer {
	if currentpage < 1 || currentpage > totalpages || around < 1 || boundaries < 1 || totalpages < 1 {
		return Footer{}
	}

	list := []string{}
	var ELLIPSIS string = "..."
	var pagesAtBeginning int = boundaries
	var pagesAtEnd int = totalpages - boundaries
	var pagesAtLeft int = currentpage - around
	var pagesAtRight int = currentpage + around

	//This loop is responsible to add the elements that fits the conditions
	for i := 1; i <= totalpages; i++ {
		if i <= pagesAtBeginning || i > pagesAtEnd || i >= pagesAtLeft && i <= pagesAtRight {
			list = append(list, strconv.Itoa(i))
		} else if list[len(list)-1] == ELLIPSIS {
			continue
		} else {
			list = append(list, ELLIPSIS)
		}
	}

	return Footer{
		List: list,
	}
}
