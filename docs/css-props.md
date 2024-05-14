# Tabber CSS custom props

Here below is the list of all CSS variables defined and their default values:

## Colors

[Source Code](../tabber-css.templ#L16-L75).

| CSS Property          | Default Value                                                                                                                                                       | Description                                     |
| :-------------------- | :------------------------------------------------------------------------------------------------------------------------------------------------------------------ | :---------------------------------------------- |
| `--gtb-color-space`   | _oklab_                                                                                                                                                             | The color space used for color mixing.          |
| `--gtb-base`          | <div style="display: flex; gap: 0.25rem;"><span style="width: 20px; height: 20px; background-color: hsl(215deg 16.3% 46.9%)"></span>_hsl(215deg 16.3% 46.9%)_</div> | Defines the base color of the tabber component. |
| `--gtb-base-10`       | Result of color-mix function                                                                                                                                        | Shade of the base color at 10% brightness.      |
| `--gtb-base-20`       | Result of color-mix function                                                                                                                                        | Shade of the base color at 20% brightness.      |
| `--gtb-base-30`       | Result of color-mix function                                                                                                                                        | Shade of the base color at 30% brightness.      |
| `--gtb-base-40`       | Result of color-mix function                                                                                                                                        | Shade of the base color at 40% brightness.      |
| `--gtb-base-50`       | _var(--gtb-base)_                                                                                                                                                   | The base color itself.                          |
| `--gtb-base-60`       | Result of color-mix function                                                                                                                                        | Shade of the base color at 60% brightness.      |
| `--gtb-base-70`       | Result of color-mix function                                                                                                                                        | Shade of the base color at 70% brightness.      |
| `--gtb-base-80`       | Result of color-mix function                                                                                                                                        | Shade of the base color at 80% brightness.      |
| `--gtb-base-90`       | Result of color-mix function                                                                                                                                        | Shade of the base color at 90% brightness.      |
| `--gtb-base-lightest` | _var(--gtb-base-10)_                                                                                                                                                | The lightest shade of the base color.           |
| `--gtb-base-lighter`  | _var(--gtb-base-20)_                                                                                                                                                | A lighter shade of the base color.              |
| `--gtb-base-light`    | _var(--gtb-base-40)_                                                                                                                                                | A light shade of the base color.                |
| `--gtb-base-dark`     | _var(--gtb-base-60)_                                                                                                                                                | A dark shade of the base color.                 |
| `--gtb-base-darker`   | _var(--gtb-base-80)_                                                                                                                                                | A darker shade of the base color.               |
| `--gtb-base-darkest`  | _var(--gtb-base-90)_                                                                                                                                                | The darkest shade of the base color.            |

## Simplified set of css custom props to make themes

| CSS Property          | Default Value                                                                                                                                           | Description                                         |
| :-------------------- | :------------------------------------------------------------------------------------------------------------------------------------------------------ | :-------------------------------------------------- |
| `--gtb-color`         | _var(--gtb-base-darker)_                                                                                                                                | The color used for text and icons in normal state.  |
| `--gtb-color-hover`   | _var(--gtb-base-darkest)_                                                                                                                               | The color used for text and icons on hover state.   |
| `--gtb-color-focus`   | _var(--gtb-base-darkest)_                                                                                                                               | The color used for text and icons on focus state.   |
| `--gtb-color-active`  | _var(--gtb-base-darkest)_                                                                                                                               | The color used for text and icons on active state.  |
| `--gtb-bg`            | <div style="display: flex; gap: 0.25rem;"><span style="width: 20px; height: 20px; background-color: hsl(0deg 0% 100%)"></span>_hsl(0deg 0% 100%)_</div> | The background color of the tabber component.       |
| `--gtb-bg-hover`      | _var(--gtb-base-lightest)_                                                                                                                              | The background color of the tabber on hover state.  |
| `--gtb-bg-focus`      | _var(--gtb-base-lightest)_                                                                                                                              | The background color of the tabber on focus state.  |
| `--gtb-bg-active`     | _var(--gtb-bg)_</div>                                                                                                                                   | The background color of the tabber on active state. |
| `--gtb-border-color`  | _var(--gtb-base-lighter)_                                                                                                                               | The border color of the tabber component.           |
| `--gtb-border-active` | _var(--gtb-base-light)_                                                                                                                                 | The border color of the tabber on active state.     |
| `--gtb-nav-bg`        | _var(--gtb-base-lighter)_                                                                                                                               | The background color of the tab list.               |

## Root

| CSS Property              | Default Value | Description                                        |
| :------------------------ | :------------ | :------------------------------------------------- |
| `--gtb-root-width`        | _100%_        | The width of the tabber component.                 |
| `--gtb-root-max-width`    | _fit-content_ | The maximum width of the tabber component.         |
| `--gtb-root-my`           | _0_           | The margin on the y-axis of the tabber component.  |
| `--gtb-root-mx`           | _0_           | The margin on the x-axis of the tabber component.  |
| `--gtb-root-py`           | _0_           | The padding on the y-axis of the tabber component. |
| `--gtb-root-px`           | _0_           | The padding on the x-axis of the tabber component. |
| `--gtb-root-border-width` | _0_           | The top border width of the tabber component.      |
| `--gtb-root-border-style` | _solid_       | The top border style of the tabber component.      |
| `--gtb-root-border-color` | _transparent_ | The top border color of the tabber component.      |
| `--gtb-root-border-raiud` | _0_           | The top border radius of the tabber component.     |

## TabList

| CSS Property                      | Default Value             | Description                                     |
| :-------------------------------- | :------------------------ | :---------------------------------------------- |
| `--gtb-tab-list-width`            | _auto_                    | The width of the tab list.                      |
| `--gtb-tab-list-max-width`        | _calc(100% + 20%)_        | The maximum width of the tab list.              |
| `--gtb-tab-list-alignment`        | _flex-start_              | The alignment of items within the tab list.     |
| `--gtb-tab-list-my`               | _0_                       | The margin on the y-axis of the tab list.       |
| `--gtb-tab-list-mx`               | _0_                       | The margin on the x-axis of the tab list.       |
| `--gtb-tab-list-py`               | _0_                       | The padding on the y-axis of the tab list.      |
| `--gtb-tab-list-px`               | _0_                       | The padding on the x-axis of the tab list.      |
| `--gtb-tab-list-gap`              | _0.5em_                   | The gap between tab list items.                 |
| `--gtb-tab-list-bg`               | _var(--gtb-bg)_           | The background color of the tab list.           |
| `--gtb-tab-list-border-t-width`   | _0_                       | The top border width of the tab list.           |
| `--gtb-tab-list-border-t-style`   | _solid_                   | The top border style of the tab list.           |
| `--gtb-tab-list-border-t-color`   | _var(--gtb-border-color)_ | The top border color of the tab list.           |
| `--gtb-tab-list-border-r-width`   | _0_                       | The right border width of the tab list.         |
| `--gtb-tab-list-border-r-style`   | _solid_                   | The right border style of the tab list.         |
| `--gtb-tab-list-border-r-color`   | _var(--gtb-border-color)_ | The right border color of the tab list.         |
| `--gtb-tab-list-border-b-width`   | _0_                       | The bottom border width of the tab list.        |
| `--gtb-tab-list-border-b-style`   | _solid_                   | The bottom border style of the tab list.        |
| `--gtb-tab-list-border-b-color`   | _var(--gtb-border-color)_ | The bottom border color of the tab list.        |
| `--gtb-tab-list-border-l-width`   | _0_                       | The left border width of the tab list.          |
| `--gtb-tab-list-border-l-style`   | _solid_                   | The left border style of the tab list.          |
| `--gtb-tab-list-border-l-color`   | _var(--gtb-border-color)_ | The left border color of the tab list.          |
| `--gtb-tab-list-border-tl-radius` | _0_                       | The top-left border radius of the tab list.     |
| `--gtb-tab-list-border-tr-radius` | _0_                       | The top-right border radius of the tab list.    |
| `--gtb-tab-list-border-bl-radius` | _0_                       | The bottom-left border radius of the tab list.  |
| `--gtb-tab-list-border-br-radius` | _0_                       | The bottom-right border radius of the tab list. |

## Tab

| CSS Property                     | Default Value                   | Description                                       |
| :------------------------------- | :------------------------------ | :------------------------------------------------ |
| `--gtb-tab-py`                   | _0.75em_                        | The padding on the y-axis of the tab item.        |
| `--gtb-tab-px`                   | _1em_                           | The padding on the x-axis of the tab item.        |
| `--gtb-tab-gap`                  | _0.375em_                       | The gap between tab items.                        |
| `--gtb-tab-color`                | _var(--gtb-color)_              | The color of the tab item.                        |
| `--gtb-tab-color-hover`          | _var(--gtb-color-hover)_        | The color of the tab item on hover.               |
| `--gtb-tab-color-focus`          | _var(--gtb-color-focus)_        | The color of the tab item on focus.               |
| `--gtb-tab-color-active`         | _var(--gtb-color-active)_       | The color of the active tab item.                 |
| `--gtb-tab-font-family`          | _inherit_                       | The font family of the tab item.                  |
| `--gtb-tab-bg`                   | _transparent_                   | The background color of the tab item.             |
| `--gtb-tab-bg-hover`             | _var(--gtb-bg-hover)_           | The background color of the tab item on hover.    |
| `--gtb-tab-bg-focus`             | _var(--gtb-bg-focus)_           | The background color of the tab item on focus.    |
| `--gtb-tab-bg-active`            | _var(--gtb-bg-active)_          | The background color of the active tab item.      |
| `--gtb-tab-border-t-width`       | _0_                             | The top border width of the tab item.             |
| `--gtb-tab-border-t-style`       | _solid_                         | The top border style of the tab item.             |
| `--gtb-tab-border-t-color`       | _var(--gtb-border-color)_       | The top border color of the tab item.             |
| `--gtb-tab-border-r-width`       | _0_                             | The right border width of the tab item.           |
| `--gtb-tab-border-r-style`       | _solid_                         | The right border style of the tab item.           |
| `--gtb-tab-border-r-color`       | _var(--gtb-border-color)_       | The right border color of the tab item.           |
| `--gtb-tab-border-b-width`       | _0_                             | The bottom border width of the tab item.          |
| `--gtb-tab-border-b-style`       | _solid_                         | The bottom border style of the tab item.          |
| `--gtb-tab-border-b-color`       | _var(--gtb-border-color)_       | The bottom border color of the tab item.          |
| `--gtb-tab-border-l-width`       | _0_                             | The left border width of the tab item.            |
| `--gtb-tab-border-l-style`       | _solid_                         | The left border style of the tab item.            |
| `--gtb-tab-border-l-color`       | _var(--gtb-border-color)_       | The left border color of the tab item.            |
| `--gtb-tab-border-t-color-hover` | _var(--gtb-tab-border-t-color)_ | The top border color of the tab item on hover.    |
| `--gtb-tab-border-r-color-hover` | _var(--gtb-tab-border-r-color)_ | The right border color of the tab item on hover.  |
| `--gtb-tab-border-b-color-hover` | _var(--gtb-tab-border-b-color)_ | The bottom border color of the tab item on hover. |
| `--gtb-tab-border-l-color-hover` | _var(--gtb-tab-border-l-color)_ | The left border color of the tab item on hover.   |
| `--gtb-tab-border-tl-radius`     | _0_                             | The top-left border radius of the tab item.       |
| `--gtb-tab-border-tr-radius`     | _0_                             | The top-right border radius of the tab item.      |
| `--gtb-tab-border-bl-radius`     | _0_                             | The bottom-left border radius of the tab item.    |
| `--gtb-tab-border-br-radius`     | _0_                             | The bottom-right border radius of the tab item.   |
| `--gtb-tab-ring-width`           | _0_                             | The width of the tab item's focus ring.           |
| `--gtb-tab-ring-style`           | _solid_                         | The style of the tab item's focus ring.           |
| `--gtb-tab-ring-color`           | _var(--gtb-base-light)_         | The color of the tab item's focus ring.           |
| `--gtb-tab-ring-offset`          | _0.125rem_                      | The offset of the tab item's focus ring.          |

## TabPanel

| CSS Property                   | Default Value             | Description                                      |
| :----------------------------- | :------------------------ | :----------------------------------------------- |
| `--gtb-panel-width`            | _100%_                    | The width of the tab panel.                      |
| `--gtb-panel-max-width`        | _680px_                   | The maximum width of the tab panel.              |
| `--gtb-panel-my`               | _0_                       | The margin on the y-axis of the tab panel.       |
| `--gtb-panel-mx`               | _0_                       | The margin on the x-axis of the tab panel.       |
| `--gtb-panel-py`               | _1.25em_                  | The padding on the y-axis of the tab panel.      |
| `--gtb-panel-px`               | _1.25em_                  | The padding on the x-axis of the tab panel.      |
| `--gtb-panel-color`            | _var(--gtb-base-darker)_  | The color of the tab panel.                      |
| `--gtb-panel-bg`               | _var(--gtb-bg)_           | The background color of the tab panel.           |
| `--gtb-panel-border-t-width`   | _0_                       | The top border width of the tab panel.           |
| `--gtb-panel-border-t-style`   | _solid_                   | The top border style of the tab panel.           |
| `--gtb-panel-border-t-color`   | _var(--gtb-border-color)_ | The top border color of the tab panel.           |
| `--gtb-panel-border-r-width`   | _0_                       | The right border width of the tab panel.         |
| `--gtb-panel-border-r-style`   | _solid_                   | The right border style of the tab panel.         |
| `--gtb-panel-border-r-color`   | _var(--gtb-border-color)_ | The right border color of the tab panel.         |
| `--gtb-panel-border-b-width`   | _0_                       | The bottom border width of the tab panel.        |
| `--gtb-panel-border-b-style`   | _solid_                   | The bottom border style of the tab panel.        |
| `--gtb-panel-border-b-color`   | _var(--gtb-border-color)_ | The bottom border color of the tab panel.        |
| `--gtb-panel-border-l-width`   | _0_                       | The left border width of the tab panel.          |
| `--gtb-panel-border-l-style`   | _solid_                   | The left border style of the tab panel.          |
| `--gtb-panel-border-l-color`   | _var(--gtb-border-color)_ | The left border color of the tab panel.          |
| `--gtb-panel-border-tl-radius` | _0_                       | The top-left border radius of the tab panel.     |
| `--gtb-panel-border-tr-radius` | _0_                       | The top-right border radius of the tab panel.    |
| `--gtb-panel-border-bl-radius` | _0_                       | The bottom-left border radius of the tab panel.  |
| `--gtb-panel-border-br-radius` | _0_                       | The bottom-right border radius of the tab panel. |
