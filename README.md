<h1 align="center" style="font-size: 2.5rem;">
  tabber
</h1>
<p align="center">
    <a href="https://github.com/indaco/tabber/blob/main/LICENSE" target="_blank">
        <img src="https://img.shields.io/badge/license-mit-blue?style=flat-square&logo=none" alt="license" />
    </a>
     &nbsp;
     <a href="https://goreportcard.com/report/github.com/indaco/tabber/" target="_blank">
        <img src="https://goreportcard.com/badge/github.com/indaco/tabber" alt="go report card" />
    </a>
    &nbsp;
    <a href="https://pkg.go.dev/github.com/indaco/tabber/" target="_blank">
        <img src="https://pkg.go.dev/badge/github.com/indaco/tabber/.svg" alt="go reference" />
    </a>
    &nbsp;
    <a href="https://www.jetify.com/devbox/docs/contributor-quickstart/">
      <img
          src="https://www.jetify.com/img/devbox/shield_moon.svg"
          alt="Built with Devbox"
      />
  </a>
</p>

<p align="center">
  <a href="#features">Features</a> •
  <a href="#installation">Installation</a> •
  <a href="#usage">Usage</a> •
  <a href="#api-reference">API Reference</a> •
  <a href="#accessibility-a11y">Accessibility (A11Y)</a> •
  <a href="#theming">Theming</a> •
  <a href="#examples">Examples</a> •
  <a href="#contributing">Contributing</a>
</p>

A fully accessible, configurable and themeable server-rendered Tab component for Go web applications. Built with [templ](https://github.com/a-h/templ) library for seamless integration with Go-based web frontends.

## Features

- **No External Dependencies**: Built with native Go and the `templ` library, `tabber` requires no external CSS or JavaScript dependencies.
- **Accessibility**: Compliant with the [WAI-ARIA Tabs Design Pattern](https://www.w3.org/WAI/ARIA/apg/patterns/tabs/), `tabber` ensures accessibility for all users, with proper use of ARIA attributes and keyboard navigation support.
- **Headless Component**: `tabber` comes as a headless component, providing flexible styling options that can be easily customized to suit your design needs.
- **Pre-defined variants**: Choose from predefined variants or define your own custom themes to match your application's style. **Users can create their own tabber style variant as a re-usable templ component and use it.**
- **Responsive Design**: Built with responsiveness in mind, `tabber` adjusts seamlessly to different screen sizes and devices, ensuring a consistent user experience across platforms.
- **Multiple Variants**: `tabber` allows for enabling multiple variants on the same page, providing versatility and adaptability to diverse design requirements.
- **Multiple tabs**: multiple `tabber` components on the same page, each with its own style.
- **Customizable Configuration**: `tabber` offers various configuration options, including positioning, alignment, and more, allowing you to tailor its behavior to fit your specific use case.
- **Auto-generated or User-defined Dark Mode**: Dark mode can be automatically generated based on the primary color defined or set by the user.
- **Themeable with CSS Variables**: Easily customize the appearance of `tabber` using CSS variables, with built-in support for theming.

![Demo](statics/demo.gif)

> **Disclaimer**: Please note that the demo video may not include the latest features and updates of the project. However, it still accurately represents the overall concepts and functionality.

## Installation

To install the `tabber` module, use the `go get` command:

```sh
go get github.com/indaco/tabber@latest
```

Ensure your project is using Go Modules (it will have a `go.mod` file in its root if it already does).

## Usage

> [!TIP]
> refer to the [Examples](#examples) section.

Import the `tabber` module into your project:

```go
import "github.com/indaco/tabber"
```

### Configuration & Context

```go
// Default options
tabberConfig := tabber.NewConfigBuilder().Build()
```

#### Available Options

Users can access each configuration option using the corresponding `With` method, such as `tabber.WithActiveTab(2)` or `tabber.WithPlacement(tabber.Left)`.

| Option      | Type          | Description                                                                                                        |
| :---------- | :------------ | :----------------------------------------------------------------------------------------------------------------- |
| `ActiveTab` | _number_      | Specifies the index of the initially active tab. The indexing starts from 1 for the first tab.                     |
| `Placement`  | [_Placement_]  | Determines the position of the tab list relative to the tab panels. Options: `Top`, `Bottom`, `Left`, and `Right`. |
| `Variants`  | [][_Variant_] | Enables specific visual variants for the tab component.                                                            |
| `Alignment` | [_Alignment_] | Defines the alignment of the tab items within the tab list. Options: `Start`, `Center`, and `End`.                 |

To allow configurations to be accessible at any level in the tabs hierarchy, `tabber` makes use of the templ component context and the implicit `ctx` variable. You can read more about templ component context [here](https://templ.guide/syntax-and-usage/context#using-context).

In your function handler, create a configuration for `tabber` and attach it to the request context passed into the handler function.

```go
func HandleHome(w http.ResponseWriter, r *http.Request) {
  tabberConfig := tabber.NewConfigBuilder().WithVariant(tabber.Underlined).Build()

  configMap := tabber.NewConfigMap()
  configMap.Add("my-tab", tabberConfig)
  ctx := context.WithValue(r.Context(), tabber.ConfigContextKey, configMap)
  err := Page().Render(ctx, w)
  if err != nil {
    return
  }
}
```

#### Variants

By default `tabber` is headless, you have to enable single or multiple predefined variants.

```go
  tabberConfig := tabber.NewConfigBuilder().WithVariant(tabber.Underlined).Build()
  // or enabling multiple variants. The first value is the default one.
  tabberConfig := tabber.NewConfigBuilder().WithVariants(tabber.Underlined, tabber.Rounded).Build()
```

Available Predefined Variants:

- `Accent`
- `Bordered`
- `Grouped`
- `Rounded`
- `Pills`
- `Underlined`

You can also define and use your own variant. Please, refer to [Create Your Own Tabber Variant](#4-create-your-own-tabber-variant).

### Tab Component Structure - Markup

In your `templ` file, defines the structure for the tab component.

> [!IMPORTANT]
> It is crucial to ensure that the value passed to `tabber.Root` **matches** the one used when adding the `gropdownConfig` to the `configMap` as per step above. This ensures that multiple tabs on the same page function independently.

```go
const (
  profileIcon = `<svg ...></svg>`
  settingsIcon = `<svg ...></svg>`
)

templ Page() {
  @tabber.Root("my-tab") {
    @tabber.TabList() {
      @tabber.Tab(1, "First", tabber.Icon(profileIcon))
      @tabber.Tab(2, "Second", tabber.Icon(settingsIcon))
      @tabber.Tab(3, "Third")
    }
    @tabber.TabPanel(1) {
      <h3>1 - Heading</h3>
      <p>Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor
          incididunt ut labore et dolore magna aliqua.</p>
    }
    @tabber.TabPanel(2) {
      <h3>2 - Heading</h3>
      <p>Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor
          incididunt ut labore et dolore magna aliqua.</p>
    }
    @tabber.TabPanel(3) {
      <h3>3 - Heading</h3>
      <p>Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor
          incididunt ut labore et dolore magna aliqua.</p>
    }
  }
}
```

### CSS and Javascript

`tabber` leverages the templ library's features, including CSS Components and JavaScript Templates, to encapsulate all necessary styling and functionality without relying on external dependencies.

- `TabberCSS`: Supplies the required CSS, encapsulating the visual design and layout specifics of the component. Moreover, `TabberCSS` selectively includes styles specific to the enabled style variant, ensuring that only essential styles are loaded to optimize page performance and minimize unnecessary overhead.
- `TabberJS`: Provides the JavaScript logic essential for dynamic behaviors such as displaying, keyboard navigation, and interaction with the component.

To facilitate integration with Go's `template/html` standard library, `tabber` includes a dedicated `HTMLGenerator` type to seamlessly integrate the component into web applications built with Go's `html/template` standard library.

There are methods acting as wrappers to the templ's `templ.ToGoHTML`, generate the necessary HTML to be embedded them into server-rendered pages:

- `TabberCSSToGoHTML`: Renders the `TabberCSS` component into a `template.HTML` value.
- `TabberJSToGoHTML`: Renders the `TabberJS` component into a `template.HTML` value.

## API Reference

<h3 style="padding: 0.25rem; width: fit-content; color: #6b7280; background-color: #bef264;">tabber.<span style="font-weight: 600; color: #1f2937;">Root</span></h3>

| Property | Type     | Description                                  |
| :------- | :------- | :------------------------------------------- |
| `id`     | _string_ | The unique identifier for the tab component. |

<h3 style="padding: 0.25rem; width: fit-content; color: #6b7280; background-color: #bef264;">tabber.<span style="font-weight: 600; color: #1f2937;">TabList</span></h3>

None

<h3 style="padding: 0.25rem; width: fit-content; color: #6b7280; background-color: #bef264;">tabber.<span style="font-weight: 600; color: #1f2937;">Tab</span></h3>

| Property | Type        | Description                                 |
| :------- | :---------- | :------------------------------------------ |
| `value`  | _int_       | The unique identifier for the tab.          |
| `label`  | _string_    | The text displayed for the tab button.      |
| `icon`   | [_TabIcon_] | The icon displayed close to the tab button. |

<h3 style="padding: 0.25rem; width: fit-content; color: #6b7280; background-color: #bef264;">tabber.<span style="font-weight: 600; color: #1f2937;">Panel</span></h3>

| Property | Type  | Description                                                           |
| :------- | :---- | :-------------------------------------------------------------------- |
| `value`  | _int_ | The unique identifier for the tab panel referenced by the tab button. |

## Accessibility (A11Y)

The tab component is designed to be accessible to screen readers and supports keyboard navigation according to the [WAI-ARIA Tabs Design Pattern](https://www.w3.org/WAI/ARIA/apg/patterns/tabs/).

### Screen Reader Support

Ensure that proper ARIA attributes are used to convey the state and role of the tab elements.

### Keyboard Interaction

- **Focusing the Tab Item**:
  - Use the `Tab` key to navigate to the tab button. Pressing `Enter` or `Space` will switch the tab.
- **Navigating within the Tab component**:
  - Use the `Arrow` keys to move focus between items within the tabs list.
  - Pressing `Home` or `End` keys will move focus to the first or last item respectively.
  - Use `A-Z` or `a-z` keys to move focus to the next tab item with a label that starts with the typed character if such a tab item exists. Otherwise, focus does not move.
- **Selecting an Item**:
  - Press `Enter` to select the currently focused item in the tabs list.

## Theming

Customizing the appearance of `tabber` to align with your design preferences is both straightforward and flexible, largely due to its extensive use of CSS custom properties (CSS variables) prefixed with `gtb`. For a comprehensive list of CSS custom properties, along with their default values and descriptions, please consult the tabber [CSS custom Props](./docs/css-props.md) document.

There are three primary methods you can employ to easily adjust `tabber` styles and variants. Depending on your specific requirements, choose the most suitable approach to achieve your desired styling.

In addition, you would create your own `tabber` variant as reusable `templ` component, register and use it. Refer to [Create your own tabber variant](#4-create-your-own-tabber-variant).

### 1. Customise Color Palette

By default, `tabber` utilizes a single `hsl` color and dynamically generates shades of it using the `color-mix` function in the `oklab` color space. To change the primary color, simply override the `--gtb-base` property with your desired value. Refer to the [colors](./docs/css-props.md#colors) section in the CSS custom props documentation for more details.

**[Example](_examples/theming/customise-color/)**

If you prefer a different color space other than `oklab`, you can customize it by overriding the `--gtb-color-space` property with your preferred value. For a comprehensive list of available color spaces, refer to the [documentation](https://developer.mozilla.org/en-US/docs/Web/CSS/color_value/color-mix).

To view the default base color shades and their derivations, consult the _tabber-css.templ_ file [here](./tabber-css.templ#L16-L89).

### 2. Customise Theme

If one of the predefined style variants suits your needs but you desire more flexibility in certain aspects, such as using different colors instead of shades of the base color, you can define values for the [Simplified set of css custom props](./docs/css-props.md#simplified-set-of-css-custom-props-to-make-themes).

**[Example](_examples/theming/customise-theme/)**

### 3. Customise CSS Variant

The majority of tabber components utilize CSS custom properties to facilitate the creation of style variants and enable extensive customization of styles.

**[Example](_examples/theming/customise-variant/)**

### 4. Create Your Own Tabber Variant

Here, we are defining and creating a new variant called `sketched` in the `sketched-variant.templ` file:

```go
/*
 * sketched-variant.templ
 */
import "github.com/indaco/tabber"

const Sketched tabber.Variant = "sketched"

// SketchedVariant is a new variant for tabber.
type SketchedVariant struct{}

// Implement TabberVariant interface
func (nv SketchedVariant) GetClassName() templ.Component {
  return SketchedVariantCss()
}

templ SketchedVariantCss() {
  <style>
    :root {
      [data-variant="sketched"] {
        --gtb-root-py: 1.5em;
        --gtb-root-px: 2em;
        --gtb-root-border-width: 4px;
        --gtb-root-border-style: solid;
        /* refer to the source code... */
      }
    }
  </style>
}
```

Use your newly create `sketched` variant:

```go
// register the Sketched variant and use it.
tabber.RegisterVariant(Sketched, SketchedVariant{})
tabberConfig := tabber.NewConfigBuilder().WithVariant(Sketched).Build()

configMap := tabber.NewConfigMap()
configMap.Add("sketched-tabber", tabberConfig)
```

**[Example](_examples/theming/variant-creation/)**

## Examples

To run the examples, make sure to compile them first by running `make examples`. Then, move to the example folder and run

```bash
cd _examples/demo/
go run .
```

- [demo](_examples/demo)
- [use with `template/html`](_examples/go-html-template)
- [multiple-tabs](_examples/multiple-tabs/)
- [customise color](_examples/theming/customise-color/)
- [customise theme](_examples/theming/customise-theme/)
- [customise variant](_examples/theming/customise-variant/)
- [variant creation](_examples/theming/variant-creation/)

## Contributing

Contributions are welcome! Feel free to open an issue or submit a pull request.

### Development Environment Setup

To set up a development environment for this repository, you can use [devbox](https://www.jetify.com/devbox) along with the provided `devbox.json` configuration file.

1. Install devbox by following the instructions in the [devbox documentation](https://www.jetify.com/devbox/docs/installing_devbox/).
2. Clone this repository to your local machine.
3. Navigate to the root directory of the cloned repository.
4. Run `devbox install` to install all packages mentioned in the `devbox.json` file.
5. Run `devbox shell` to start a new shell with access to the environment.
6. Once the devbox environment is set up, you can start developing, testing, and contributing to the repository.

### Using Makefile

Additionally, you can make use of the provided `Makefile` to run various tasks:

```bash
make build       # The main build target
make examples    # Process templ files in the _examples folder
make templ       # Process TEMPL files
make test        # Run go tests
make help        # Print this help message
```

## License

This project is licensed under the MIT License - see the LICENSE file for details.

<!-- Resources -->

[_Placement_]: ./types.go#L8-L11
[_Alignment_]: ./types.go#L16-L18
[_Variant_]: ./types.go#L23-29
[_TabIcon_]: ./types.go#L35-L38
