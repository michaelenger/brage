# Brage

_Brage is the Norwegian name for the ancient norse god [Bragi](https://en.wikipedia.org/wiki/Bragi),
the skaldic god of poetry._

Brage is a simple static site generator written in [Go](https://go.dev/). It
supports building pages and posts using [mustache](https://mustache.github.io/)
and [Markdown](https://www.markdownguide.org/) templates.

## Usage

Usage is based on three main commands, `init`, `serve`, and `build`, all of  which are built to work on a single source directory.

### Init

```shell
brage init [PATH]
```

`init` is used to initialise a new site and will create a bunch of files that can be used as a template when creating a new site. If no `PATH` is specified then it will generate the files in the current directory.

#### Options

* `-f, --force` Force the creation of the site contents, overwriting any existing files

### Serve

```shell
brage serve [PATH]
```

`serve` will serve the site specified in the `PATH` (or the current directory if nothing is specified) on port `8080`. This can be used when developing or debugging the site.

#### Options

* `-p, --port port` Port to serve the site on (default: `8080`)

### Build

```shell
brage build [PATH]
```

`build` builds the site, generating all the static HTML files and copying any assets to the appropriate location. It will read the site from the `PATH` location (or the current directory if nothing is specified) and store the generated files in a `build` subdirectory if no output path is specified.

#### Options

* `-o, --output` Path to output the site to
* `-c, --clean` Override the output assets directory, removing anything already in there

## Building Sites

Sites are defined with a config [YAML](https://yaml.org/) file, an optional layout template, one or more page templates, one or more post templates, partial templates, and optional assets.

### Config

The config for a site is specified in a `config.yaml` file in the site's directory. It can contain the following fields:

* `title` The site title
* `description` Site description
* `image` Favicon
* `root_url` The root URL of the site
* `redirects` Map of URIs that should redirect to other URLs
* `data` A map containing any optional data you want to use in the templates

The contents of the config file is available in the templates under the `site` variable, and anything defined in the `data` field is available under `data`:

```gohtml
Welcome to {{ site.title }}.

Here is my dog: {{ data.dog }}
```

### Templates

The HTML templates are all parsed as standard [Mustache templates](https://mustache.github.io/) and HTML is not escaped, so you are forewarned that the rendering isn't going to sanitise anything for you.

#### Variables

The following variables are passed into the template and are available:

##### Site

Contains site data as defined in the `config.yaml` file:

* `site.title` Site title
* `site.description` Site description
* `site.image` Favicon/social media image
* `site.root_url` Root URL
* `site.redirects` Redirect map
* `site.posts` A list of all available posts (with their respective `path`, `title`, and `date`s)

##### Page

Contains information about the current page:

* `page.path` Path to the page
* `page.template` Contents of the page template
* `page.title` Automatically inferred title based on the path

The title for the root path is `"Home"`

##### Post

Contains information about the current post:

* `page.path` Path to the page
* `page.template` Contents of the page template
* `page.title` Automatically inferred title based on the path
* `page.date` Date of the post (as specified in the metadata)

##### Data

The `data` variable contains all the variables which were added in the `data` field in the `config.yaml` file. For example, the following config:

```yaml
data:
  bananas:
    - ripe
    - green
    - mouldy
  explosions: "all over the place"
  best_numbers:
    - name: one
      value: 1
    - name: four
      value: 4
    - name: nine
      value: 9
```

Would result in the following variables being present:

* `data.bananas` An array of strings
* `data.explosions` The string "all over the place"
* `data.best_numbers` An array of the best numbers containing maps

##### Content

In the `layout.html` file you can also use the special command ```{{{content}}}``` to output the contents of the current page.

### Layout

The `layout.html` file defines the layout of the site and is used to wrap all pages. When a page is generated its contents are stored and made available in the `content` template variable.

An example layout which doesn't add more than the site title would be as follows:

```gohtml
<head>
    <title>{{ page.title }}</title>
</head>
<body>
    {{{ content }}}
</body>
```

Using a layout template is optional, but _highly recommended_.

#### Page and Post Layout

You can use custom layouts for posts and pages by providing the `layout-page.html` or `layout-post.html` files.

### Pages

Pages are built based on template files in a `pages` subdirectory and need to have the `.html` or `.markdown` file extension for Go template and Markdown templates respectively. The URI for the page is based on its name (and subdirectory) except for any template named `index` which will have no name.

* `/pages/index.html` => `/`
* `/pages/another-page.markdown` => `/another-page`
* `/pages/sub/index.html` => `/sub`
* `/pages/sub/sub/page.html` => `/sub/sub/page`

### Posts

Posts (similar to pages) are in template files in a `posts` subdirectory and can be both HTML or Markdown, defined by their file extension. The URI for the post is also on its file name so note that there is nothing stopping you from creating a post and a page which override each other.

* `/posts/first-post.markdown` => `/first-post`
* `/posts/blog/this-is-a-subdir.markdown` => `/blog/this-is-a-subdir`
* `/posts/blog/sub/post.html` => `/blog/sub/post`

#### Post Metadata

Posts written in Markdown can define the title and date for the post in a YAML "front matter" section:

```markdown
---
title: Post title goes here
date: 2009-08-07
---
This is the actual post.
```

Dates must be defined in the `YYYY-MM-DD` format and the list of posts will be sorted to show the latest ones first.

### Partials

Any files present in the `partials` subdirectory will be available using their name with the partial syntax:

_/partials/extra.markdown_
```markdown
This is in the template.
```

_/pages/index.html_
```gohtml
This is in the page.

{{> extra }}
```

### Assets

Assets are files in the `assets` subdirectory and are copied directly to an `assets` subdirectory in the target path when building the site.

## Building

To build a binary that can work as a part of a GitHub Actions pipeline you need to run the following command:

```shell
GOARCH=amd64 GOOS=linux go build
```

## TODO

* Generate an RSS/Atom file.
* Add support for a post description and image (for meta tags).
* Provide lists of top 5/10 posts (or use a lambda?)
* Add word count and reading time to the posts.
* Warn when a post is going to override another one.
* Allow for customizing the posts prefix.
* Support tags for posts?
* Hide posts with dates in the future.
* Add filters/lambdas to make working with Mustache a bit better (date formatter, list limiter, etc.)
