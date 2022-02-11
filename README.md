# Brage

_Brage is the Norwegian name for the ancient norse god [Bragi](https://en.wikipedia.org/wiki/Bragi),
the skaldic god of poetry._

Brage is a simple static site generator written in [Go](https://go.dev/). It
supports building pages using [Go templates](https://pkg.go.dev/text/template) and
[Markdown](https://www.markdownguide.org/).

## Usage

Usage is based on three main commands, `init`, `serve`, and `build`, all of which
are built to work on a single source directory.

### Init

```shell
brage init [PATH]
```

`init` is used to initialise a new site and will create a bunch of files that can
be used as a template when creating a new site. If no `PATH` is specified then it
will generate the files in the current directory.

#### Options

* `-f, --force` Force the creation of the site contents, overwriting any existing files

### Serve

```shell
brage serve [PATH]
```

`serve` will serve the site specified in the `PATH` (or the current directory if
nothing is specified) on port `8080`. This can be used when developing or debugging
the site.

#### Options

* `-p, --port port` Port to serve the site on (default: `8080`)

### Build

```shell
brage build [PATH]
```

`build` builds the site, generating all the static HTML files and copying any assets
to the appropriate location. It will read the site from the `PATH` location (or the
current directory if nothing is specified) and store the generated files in a `build`
subdirectory if no output path is specified.

#### Options

* `-o, --output` Path to output the site to

## Building Sites

Sites are defined with a config [YAML](https://yaml.org/) file, a layout template,
one or more page templates, and optional assets.

### Config

The config for a site is specified in a `config.yaml` file in the site's directory.
It can contain the following fields:

* `title` The site title
* `description` Site description
* `image` Favicon
* `rootUrl` The root URL of the site
* `data` A map containing any optional data you want to use in the templates

The contents of the config file is available in the templates under the `.Site` variable,
and anything defined in the `data` field is available under `.Data`:

```gohtml
Welcome to {{ .Site.Title }}.

Here is my dog: {{ .Data.dog }}
```

### Layout

The `layout.html` file defines the layout of the site and is used to wrap all pages.
When a page is generated its contents are stored and made available in the `.Content`
template variable.

An example layout which doesn't add more than the site title would be as follows:

```gohtml
<head>
    <title>{{ .Page.Title }}</title>
</head>
<body>
    {{ .Content }}
</body>
```

### Pages

Pages are built based on template files in a `pages` subdirectory and need to have
the `.html` or `.markdown` file extension for Go template and Markdown templates
respectively. The URL for the page is based on its name (and subdirectory) except
for any template named `index` which will have no name.

* `/pages/index.html` => `/`
* `/pages/another-page.markdown` => `/another-page`
* `/pages/sub/index.html` => `/sub`
* `/pages/sub/sub/page.html` => `/sub/sub/page`

### Extra Templates

Any files present in the `templates` subdirectory will be available using their name
in any page templates using the `template` function. Their name is their path relative
to the `templates` directory without the file extension.

_/templates/extra.markdown_
```markdown
This is in the template.
```

_/pages/index.html_
```gohtml
This is in the page.

{{ template "extra" }}
```

### Assets

Assets are files in the `assets` subdirectory and are copied directly to an `assets`
subdirectory in the target path when building the site.

## TODO

Potential changes to the tool:

* Remove any required items from the config file and just let the whole thing be in `.Data`?
* Support not using a layout template?
* Customise the path to the assets directory?
