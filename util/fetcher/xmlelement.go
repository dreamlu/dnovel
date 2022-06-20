package fetcher

import (
	"fmt"
	"github.com/antchfx/htmlquery"
	"github.com/gocolly/colly/v2"
	"golang.org/x/net/html"
	"net/url"
	"strings"
)

func NewXMLElement(doc *colly.XMLElement) *XMLElement {
	return &XMLElement{doc}
}

type XMLElement struct {
	*colly.XMLElement
}

func (h *XMLElement) ChildText(xpathQuery string) string {
	if xpathQuery == "" {
		return ""
	}
	return h.XMLElement.ChildText(xpathQuery)
}

func (h *XMLElement) ChildAttr(xpathQuery string, attrName string) string {
	if xpathQuery == "" {
		return ""
	}
	return h.XMLElement.ChildAttr(xpathQuery, attrName)
}

func (h *XMLElement) ChildHtml(xpathQuery string) string {
	if xpathQuery == "" {
		return ""
	}
	child := htmlquery.FindOne(h.DOM.(*html.Node), xpathQuery)
	if child == nil {
		return ""
	}
	return strings.TrimSpace(htmlquery.OutputHTML(child, false))
}

func (h *XMLElement) ChildRemoveHtml(xpathQuery, remove string) string {
	if xpathQuery == "" {
		return ""
	}
	child := htmlquery.FindOne(h.DOM.(*html.Node), xpathQuery)
	if child == nil {
		return ""
	}
	child.RemoveChild(htmlquery.FindOne(h.DOM.(*html.Node), fmt.Sprintf("%s/%s", xpathQuery, remove)))
	return strings.TrimSpace(htmlquery.OutputHTML(child, false))
}

func (h *XMLElement) ChildUrl(xpathQuery string, attrName string) string {
	if xpathQuery == "" {
		return ""
	}
	href := h.ChildAttr(xpathQuery, attrName)
	uri, err := url.Parse(href)
	if href == "" || err != nil {
		return ""
	}
	baseUri, _ := url.Parse(h.Request.URL.String())
	if !uri.IsAbs() {
		return baseUri.ResolveReference(uri).String()
	} else {
		return uri.String()
	}
}

func (h *XMLElement) ChildUrlText(xpathQuery string) string {
	if xpathQuery == "" {
		return ""
	}
	href := h.ChildText(xpathQuery)
	uri, err := url.Parse(href)
	if href == "" || err != nil {
		return ""
	}
	baseUri, _ := url.Parse(h.Request.URL.String())
	if !uri.IsAbs() {
		return baseUri.ResolveReference(uri).String()
	} else {
		return uri.String()
	}
}
