// Code generated by templ - DO NOT EDIT.

// templ: version: v0.3.857
package tabber

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import templruntime "github.com/a-h/templ/runtime"

func TabberJS(configMap *ConfigMap) templ.ComponentScript {
	return templ.ComponentScript{
		Name: `__templ_TabberJS_440d`,
		Function: `function __templ_TabberJS_440d(configMap){/**
   * A utility class with helper methods.
   */
  class Helpers {
    /**
     * Checks if a given string is a single character.
     * @param {string} txt - The string to check.
     * @returns {boolean} Returns true if the string is a single character, otherwise false.
     */
    static isChar(txt) {
      const charsRegex = /\S/;
      return txt.length === 1 && charsRegex.test(txt);
    }

    /**
     * Extracts the first character from a text node.
     * @param {Node} item - The node from which to extract the first character.
     * @returns {string|null} The first character of the text node, or null if not found.
     */
    static extractFirstChar(item) {
      let firstChar = null;

      // traverse through child nodes to find text content
      const traverseChildNodes = (node) => {
        if (node.nodeType === Node.TEXT_NODE) {
          const textContent = node.textContent.trim();
          if (textContent) {
            firstChar = textContent[0].toLowerCase();
          }
        } else if (node.hasChildNodes()) {
          node.childNodes.forEach(traverseChildNodes);
        }
      };

      // start traversing from the button element
      traverseChildNodes(item);

      return firstChar;
    }
  }

  /**
   * The ` + "`" + `ComponentFocusManager` + "`" + ` class provides focus management for components with multiple
   * interactive items. It enables navigation between items, setting focus to specific items, and
   * handling keyboard interactions.
   *
   * Use it to ensure keyboard accessibility as per [WAI ARIA Patterns](https://www.w3.org/WAI/ARIA/apg/patterns/)
   * and improve user experience in menus, dropdowns, and other interactive components.
   */
  class ComponentFocusManager {
    /**
     * Constructs an instance of ComponentFocusManager.
     * @param {string} componentId - The ID of the component.
     */
    constructor(componentId) {
      this._id = componentId;
      this._items = {};
      this._firstChars = {};
      this._firstItem = {};
      this._lastItem = {};
      this._applyDOMChangesFn = async () => {
        // Default implementation: no pending state changes, resolves immediately
        return Promise.resolve();
      };

      this._items[componentId] = [];
      this._firstChars[componentId] = [];
      this._firstItem[componentId] = null;
      this._lastItem[componentId] = null;
    }

    /**
     * Gets or sets the collection of items.
     * @type {Object<string, HTMLElement[]>}
     */
    get items() {
      const { _items } = this;
      return _items;
    }

    set items(value) {
      this._items = value;
    }

    /**
     * Gets or sets the collection of first characters.
     * @type {Object<string, string[]>}
     */
    get firstChars() {
      const { _firstChars } = this;
      return _firstChars;
    }

    set firstChars(value) {
      this._firstChars = value;
    }

    /**
     * Gets or sets the first item in each collection.
     * @type {Object<string, HTMLElement | null>}
     */
    get firstItem() {
      const { _firstItem } = this;
      return _firstItem;
    }

    set firstItem(value) {
      this._firstItem = value;
    }

    /**
     * Gets or sets the last item in each collection.
     * @type {Object<string, HTMLElement | null>}
     */
    get lastItem() {
      const { _lastItem } = this;
      return _lastItem;
    }

    set lastItem(value) {
      this._lastItem = value;
    }

    /**
     * Gets or sets the applyDOMChanges promise.
     * @type {() => Promise<void>}
     */
    get applyDOMChangesFn() {
      return this._applyDOMChangesFn;
    }

    set applyDOMChangesFn(value) {
      this._applyDOMChangesFn = value;
    }

    /**
     * Runs the focus manager to initialize the collections.
     */
    run() {
      this._items[this._id].forEach((item) => {
        const menuItemContent = item.textContent?.trim().toLowerCase()[0];
        if (menuItemContent) this._firstChars[this._id].push(menuItemContent);

        if (!this._firstItem[this._id]) {
          this._firstItem[this._id] = item;
        }
        this._lastItem[this._id] = item;
      });
    }

    /**
     * Sets the focus to the first item.
     * @param {string} [cId] - Optional ID of the component.
     */
    setFocusToFirstItem(cId) {
      const id = this._resolveId(cId);
      this._setFocusToItem(this._firstItem[id]);
    }

    /**
     * Sets the focus to the last item.
     * @param {string} [cId] - Optional ID of the component.
     */
    setFocusToLastItem(cId) {
      const id = this._resolveId(cId);
      this._setFocusToItem(this._lastItem[id]);
    }

    /**
     * Sets the focus to the previous item relative to the current item.
     * @param {HTMLElement} currentItem - The current item.
     * @param {string} [cId] - Optional ID of the component.
     * @returns {HTMLElement} The new focused item.
     */
    setFocusToPreviousItem(currentItem, cId) {
      const id = this._resolveId(cId);
      let newMenuItem, index;

      if (currentItem === this._firstItem[id]) {
        newMenuItem = this._lastItem[id];
      } else {
        index = this._items[id].indexOf(currentItem);
        newMenuItem = this._items[id][index - 1];
      }

      this._setFocusToItem(newMenuItem);
      return newMenuItem;
    }

    /**
     * Sets the focus to the next item relative to the current item.
     * @param {HTMLElement} currentItem - The current item.
     * @param {string} [cId] - Optional ID of the component.
     * @returns {HTMLElement} The new focused item.
     */
    setFocusToNextItem(currentItem, cId) {
      const id = this._resolveId(cId);
      let newMenuItem, index;

      if (currentItem === this._lastItem[id]) {
        newMenuItem = this._firstItem[id];
      } else {
        index = this._items[id].indexOf(currentItem);
        newMenuItem = this._items[id][index + 1];
      }

      this._setFocusToItem(newMenuItem);
      return newMenuItem;
    }

    /**
     * Sets the focus to the item whose panel starts with the specified character.
     * @param {HTMLElement} currentItem - The current item.
     * @param {string} c - The character to match.
     */
    setFocusByFirstChar(currentItem, c) {
      let start, index;

      if (c.length > 1) return;
      c = c.toLowerCase();
      start = this._items[this._id].indexOf(currentItem) + 1;
      if (start >= this._items[this._id].length) {
        start = 0;
      }
      index = this._firstChars[this._id].indexOf(c, start);
      if (index === -1) {
        index = this._firstChars[this._id].indexOf(c, 0);
      }
      if (index > -1) {
        this._setFocusToItem(this._items[this._id][index]);
      }
    }

    /**
     * Checks if the given node is a submenu.
     * @param {HTMLElement} node - The node to check.
     * @returns {boolean} A boolean indicating whether the node is a submenu.
     */
    isSubMenu(node) {
      return node.getAttribute("aria-haspopup") === "true";
    }

    /**
     * Check if a value is null or undefined.
     * @param {string} [value] - The value to check.
     * @returns {boolean} A boolean indicating whether the value is nullish
     * @private
     */
    _isNullish(value) {
      return value === null || value === undefined;
    }

    /**
     * Resolves the ID to use based on the provided ID or the default ID of the component.
     * @param {string} [id] - The optional ID to resolve.
     * @returns {string} The resolved ID.
     * @private
     */
    _resolveId(id) {
      return !this._isNullish(id) ? id : this._id;
    }

    /**
     * Sets the focus to the given item.
     * @param {HTMLElement | null} item - The item to set focus to.
     * @private
     */
    async _setFocusToItem(item) {
      if (this._items[this._id] && item) {
        await Promise.all(
          this._items[this._id].map(async (itemNode) => {
            if (itemNode === item) {
              itemNode.tabIndex = 0;
              await this.applyDOMChangesFn();
              itemNode.focus();
            } else {
              itemNode.tabIndex = -1;
            }
          }),
        );
      }
    }
  }

  /**
   * Utility class for managing attributes of DOM nodes.
   */
  class NodeAttributesManager {
    /**
     * Sets attributes on the selected tab button.
     * @param {HTMLElement} selectedTabNode - The selected tab button element.
     */
    static setActiveAttributes(selectedTabNode) {
      selectedTabNode.classList.add("active");
      selectedTabNode.setAttribute("aria-selected", "true");
      selectedTabNode.setAttribute("tabindex", "0");
    }

    /**
     * Removes attributes from all tab buttons.
     * @param {NodeList} tabButtons - The NodeList containing tab buttons.
     */
    static resetAttributes(tabButtons) {
      tabButtons.forEach((btn) => {
        btn.classList.remove("active");
        btn.setAttribute("aria-selected", "false");
        btn.setAttribute("tabindex", "-1");
      });
    }

    /**
     * Sets attributes on the selected tab button and removes attributes from others.
     * @param {string} selectedDataTabId - The ID of the selected tab.
     * @param {NodeList} tabButtons - The NodeList containing tab buttons.
     */
    static updateAttributes(node, selectedDataTabId, tabButtons) {
      NodeAttributesManager.resetAttributes(tabButtons);
      const selectedTabNode = node.querySelector(` + "`" + `#tab-${selectedDataTabId}-btn` + "`" + `);
      if (selectedTabNode) {
        NodeAttributesManager.setActiveAttributes(selectedTabNode);
      }
    }

    /**
     * Adds or removes attributes to make the specified tab button active or inactive.
     * @param {HTMLElement} tabButton - The tab button element.
     * @param {boolean} isActive - Whether the tab button should be active.
     * @deprecated
     */
    static toggleTabButtonAttributes(tabButton, isActive) {
      if (isActive) {
        tabButton.classList.add("active");
        tabButton.setAttribute("aria-selected", "true");
        tabButton.setAttribute("tabindex", "0");
      } else {
        tabButton.classList.remove("active");
        tabButton.setAttribute("aria-selected", "false");
        tabButton.setAttribute("tabindex", "-1");
      }
    }

    /**
     * Sets the ARIA orientation attribute of a tab list element.
     * @param {HTMLElement} tabListEl - The tab list element.
     * @param {string} orientation - The orientation value ("horizontal" or "vertical").
     */
    static setAriaOrientation(tabListEl, orientation) {
      tabListEl.setAttribute("aria-orientation", orientation);
    }
  }

  /**
   * A utility class for managing animations based on variant types.
   */
  class AnimationManager {
    /**
     * The variant type for underlined animation.
     * @type {string}
     */
    static UNDERLINED_VARIANT = "underlined";

    /**
     * The variant type for pills animation.
     * @type {string}
     */
    static PILLS_VARIANT = "pills";

    /**
     * An array containing all variants with animations.
     * @type {string[]}
     */
    static variantsWithAnimations = [
      AnimationManager.UNDERLINED_VARIANT,
      AnimationManager.PILLS_VARIANT,
    ];

    /**
     * Checks if a variant has animation.
     * @param {string} variant - The variant type to check.
     * @returns {boolean} Returns true if the variant has animation, otherwise false.
     */
    static isVariantWithAnimation(variant) {
      return AnimationManager.variantsWithAnimations.includes(variant);
    }

    /**
     * Checks if the window is at a given breakpoint.
     * @param {string} breakpoint - The breakpoint media query.
     * @returns {boolean} Returns true if the window matches the breakpoint, otherwise false.
     */
    static isWindowAtBreakpoint(breakpoint) {
      return window.matchMedia(breakpoint).matches;
    }

    /**
     * Checks if the window matches the animation breakpoint for a specific variant.
     * @param {string} variant - The variant type to check.
     * @returns {boolean} Returns true if the window matches the animation breakpoint for the variant, otherwise false.
     */
    static windowMediaMatchesAnimationBreakpoint(variant) {
      switch (variant) {
        case AnimationManager.UNDERLINED_VARIANT:
          return AnimationManager.isWindowAtBreakpoint("(min-width: 480px)");
        case AnimationManager.PILLS_VARIANT:
          return AnimationManager.isWindowAtBreakpoint("(min-width: 650px)");
        default:
          return false;
      }
    }

    /**
     * Handles animation for a target element based on its variant and window size.
     * @param {HTMLElement} target - The target element to apply animation to.
     * @param {number} [animationDelay=100] - The delay before applying the animation.
     */
    static handle(target, animationDelay = 100) {
      const currentVariant = target.getAttribute("data-variant");

      if (
        AnimationManager.isVariantWithAnimation(currentVariant) &&
        AnimationManager.windowMediaMatchesAnimationBreakpoint(currentVariant)
      ) {
        const activeTabEl = target.querySelector('[aria-selected="true"]');
        setTimeout(() => {
          target.style.setProperty("--_left", ` + "`" + `${activeTabEl.offsetLeft}px` + "`" + `);
          target.style.setProperty("--_width", ` + "`" + `${activeTabEl.offsetWidth}px` + "`" + `);
        }, animationDelay);
      }
    }
  }

  /**
   * Manages accessibility (a11y) features for tab components.
   */
  class A11YManager {
    /**
     * Constructs a new A11YManager instance.
     * @param {HTMLElement} node - The container node for the tab component.
     * @param {object} options - Options for configuring the tab component.
     * @param {string} options.activeTab - The ID of the initially active tab.
     * @param {string} options.variant - The variant of the tab component.
     */
    constructor(node, options) {
      this.node = node;
      this.options = options;
      this.activeTab = options?.activeTab;
      this.variant = options?.variant;
      this.componentId = node.getAttribute("id");
      this.focusManager = new ComponentFocusManager(this.componentId);
      this.tabList = {};
      this.firstChars = {};
      this.firstTab = null;
      this.lastTab = null;
    }

    /**
     * Initializes tab lists and first characters.
    */
    initialize() {
      this.tabList[this.componentId] = [];
      this.firstChars[this.componentId] = [];
    }

    /**
     * Sets the active tab.
     * @param {string} selectedDataTabId - The ID of the selected tab.
     */
    setActiveTab(selectedDataTabId) {
      this.activeTab = String(
        this.tabList[this.componentId].indexOf(selectedDataTabId) + 1,
      );
    }

    /**
     * Switches the active tab.
     * @param {string} currentTabId - The ID of the current tab.
     */
    switchActiveTab(currentTabId) {
      this.variant = this.node.getAttribute("data-variant");
      const currentTabEl = this.node.querySelector(` + "`" + `#tab-${currentTabId}-btn` + "`" + `);

      if (currentTabEl) {
        NodeAttributesManager.updateAttributes(
          this.node,
          currentTabId,
          this.tabButtons,
        );

        currentTabEl.classList.add("active");
        currentTabEl.setAttribute("aria-selected", "true");
        currentTabEl.setAttribute("tabindex", "0");

        if (
          AnimationManager.isVariantWithAnimation(this.variant) &&
          AnimationManager.windowMediaMatchesAnimationBreakpoint(this.variant)
        ) {
          this.updateTabIndicator();
        }
      }
    }

    /**
     * Updates the tab indicator.
     */
    updateTabIndicator() {
      const activeTabEl = this.node.querySelector('[aria-selected="true"]');
      setTimeout(() => {
        this.node.style.setProperty("--_left", ` + "`" + `${activeTabEl.offsetLeft}px` + "`" + `);
        this.node.style.setProperty("--_width", ` + "`" + `${activeTabEl.offsetWidth}px` + "`" + `);
      }, 100);
    }

    /**
     * Switches the active tab panel.
     * @param {string} selectedDataTabId - The ID of the selected tab.
     */
    switchActiveTabPanel(selectedDataTabId) {
      const tabPanels = Array.from(this.node.querySelectorAll(".gtb_tab_panel"));
      tabPanels.forEach((panel) => {
        const id = panel.getAttribute("id");
        panel.setAttribute(
          "aria-hidden",
          id.includes(selectedDataTabId.toString()) ? "false" : "true",
        );
      });
    }

    /**
     * Handles tab events.
     * @param {Event} event - The event object.
     */
    handleTabEvent(event) {
      const targetButton = event.target.closest("button");
      if (targetButton) {
        if (event.type === "click") {
          this.onTabClick(targetButton);
        } else if (event.type === "keydown") {
          this.onKeydown(targetButton, event);
        }
      }
    }

    /**
     * Handles tab click events.
     * @param {HTMLElement} currentTab - The clicked tab element.
     */
    onTabClick(currentTab) {
      const currentDataTabId =
        currentTab.getAttribute("data-tab-id") || this.activeTab;
      this.setActiveTab(currentDataTabId);
      this.switchActiveTab(currentDataTabId);
      this.switchActiveTabPanel(currentDataTabId);
    }

    /**
     * Handles keydown events.
     * @param {HTMLElement} target - The target element.
     * @param {KeyboardEvent} event - The keyboard event.
     */
    onKeydown(target, event) {
      const selectedTabId = target.getAttribute("data-tab-id") || this.activeTab;

      if (event.shiftKey) {
        if (Helpers.isChar(event.key)) {
          this.focusManager.setFocusByFirstChar(target, event.key);
        }
      } else {
        switch (event.code) {
          case "Enter":
          case "Space":
            this.setActiveTab(selectedTabId);
            this.switchActiveTab(selectedTabId);
            this.switchActiveTabPanel(selectedTabId);
            break;
          case "Esc":
          case "Escape":
            target.blur();
            break;
          case "Left":
          case "ArrowLeft":
            this.focusManager.setFocusToPreviousItem(target);
            break;
          case "Tab":
          case "Right":
          case "ArrowRight":
            this.focusManager.setFocusToNextItem(target);
            break;
          case "Home":
          case "PageUp":
            break;
          case "End":
          case "PageDown":
            break;
          default:
            if (Helpers.isChar(event.key)) {
              this.focusManager.setFocusByFirstChar(target, event.key);
            }
            break;
        }
      }
    }

    /**
     * Sets up event listeners for tab buttons.
     */
    setupEventListeners() {
      const tabListEl = this.node.querySelector(".gtb_tab_list");
      this.tabButtons = Array.from(tabListEl.querySelectorAll("button"));
      tabListEl.addEventListener("click", (event) => this.handleTabEvent(event));
      tabListEl.addEventListener("keydown", (event) =>
        this.handleTabEvent(event),
      );
    }

    /**
     * Sets up tab panels.
     */
    setupTabPanels() {
      this.tabPanels = Array.from(this.node.querySelectorAll(".gtb_tab_panel"));
      this.tabPanels.forEach((panel) => {
        const id = panel.getAttribute("id");
        panel.setAttribute(
          "aria-hidden",
          id.includes(this.options?.activeTab.toString()) ? "false" : "true",
        );
      });
    }

    /**
     * Sets up the tab component.
     */
    setup() {
      this.initialize();
      this.setupEventListeners();
      this.setupTabPanels();
      this.focusManager.items = this.tabList;
      this.focusManager.firstChars = this.firstChars;
      this.focusManager.firstItem = this.firstTab;
      this.focusManager.lastItem = this.lastTab;
      if (AnimationManager.isVariantWithAnimation(this.variant)) {
        this.updateTabIndicator();
      }
    }
  }

  /* ************************************************************ */

  document.addEventListener("DOMContentLoaded", function () {
    // Access the value of a CSS custom properties defined in :root
    const rootStyles = window.getComputedStyle(document.documentElement);

    const orientationMap = {
      top: "vertical",
      bottom: "vertical",
      left: "horizontal",
      right: "horizontal",
    };

    function getOrientation(placement) {
      return orientationMap[placement] || "unknown_aria_orientation";
    }

    function nodeNotFoundErrorMsg(tabKey) {
      console.error(` + "`" + `Container node not found for key ${tabKey}` + "`" + `);
    }

    /**
     * Sets the background color of the specified target element based on the computed
     * body background color, with fallback to white if the body background color is transparent.
     * If a custom background color is defined in CSS, it overrides the body background color.
     * @param {CSSStyleDeclaration} rootStyles - Computed styles for the :root element.
     * @param {HTMLElement} target - The target element to set the background color for.
     */
    function setBodyBackgroundColor(rootStyles, target) {
      const bodyStyle = window.getComputedStyle(document.body);
      let backgroundColor = bodyStyle.backgroundColor;

      // Fallback to white if the body background color is transparent
      if (backgroundColor === "rgba(0, 0, 0, 0)") {
        backgroundColor = "rgb(255, 255, 255)";
      }

      // Check if a custom background color is defined in CSS
      const tabPanelBg = rootStyles.getPropertyValue('--gtb-panel-bg');
      if (tabPanelBg !== "") {
        backgroundColor = tabPanelBg;
      }

      // Set the background color property on the target element
      target.style.setProperty("--_body-bg", backgroundColor);
    }

    /**
     * Handles changes in preferred color scheme and updates the background color
     * of the specified target element accordingly.
     * @param {CSSStyleDeclaration} rootStyles - Computed styles for the :root element.
     * @param {HTMLElement} target - The target element to update the background color for.
     */
    function handleColorSchemeChange(rootStyles, target) {
      const mql = window.matchMedia("(prefers-color-scheme: dark)");
      mql.addEventListener("change", () => {
        setBodyBackgroundColor(rootStyles, target);
      });
    }

    /**
     * Handles mutations related to variant changes and updates the background color
     * of the target element accordingly.
     * @param {CSSStyleDeclaration} rootStyles - Computed styles for the :root element.
     * @param {object} tabberConfig - The configuration object for the tabber.
     * @param {MutationRecord[]} mutationsList - The list of mutations to handle.
     */
    function handleVariantChange(rootStyles, tabberConfig, mutationsList) {
      mutationsList.forEach((mutation) => {
        const target = mutation.target;
        if (!target) {
          nodeNotFoundErrorMsg(key);
          return; // Halt execution here
        }

        setBodyBackgroundColor(rootStyles, target);

        if (
          mutation.type === "attributes" &&
          mutation.attributeName === "data-variant"
        ) {
          AnimationManager.handle(target);
        }
      });
    }

    /**
     * Handles mutations related to changes in the body background color and updates
     * the background color of the target element accordingly.
     * @param {CSSStyleDeclaration} rootStyles - Computed styles for the :root element.
     * @param {object} tabberConfig - The configuration object for the tabber.
     * @param {MutationRecord[]} mutationsList - The list of mutations to handle.
     */
    function handleBodyBackgroundColorChange(rootStyles, tabberConfig, mutationsList) {
      mutationsList.forEach((mutation) => {
        const target = mutation.target;
        if (!target) {
          nodeNotFoundErrorMsg(key);
          return; // Halt execution here
        }

        if (
          mutation.type === "attributes" &&
          (mutation.attributeName === "style" ||
            mutation.attributeName === "class")
        ) {
          setBodyBackgroundColor(rootStyles, target);
        }
      });
    }

    /**
     * Handles window resize events and triggers animation handling for the specified tab.
     * @param {string} tabKey - The key identifier of the tab.
     */
    function handleWindowResize(tabKey) {
      const target = document.getElementById(` + "`" + `tabber-container-${tabKey}` + "`" + `);
      if (!target) {
        nodeNotFoundErrorMsg(tabKey);
        return; // Halt execution here
      }
      AnimationManager.handle(target);
    }

    /**
     * Initializes a tabber component with the specified configuration and container node.
     * @param {CSSStyleDeclaration} rootStyles - Computed styles for the :root element.
     * @param {tabKey} string - the id of the tabber instance.
     * @param {object} tabConfig - The configuration object for the tabber instance.
     * @param {HTMLElement} tabNode - The container node of the tabber instance.
     */
    function initializeTabberInstance(rootStyles, tabKey, tabConfig, tabNode) {
      // Initial check for preferred color scheme
      handleColorSchemeChange(rootStyles, tabNode);
      setBodyBackgroundColor(rootStyles, tabNode);

      const tabListEl = tabNode.querySelector(".gtb_tab_list");
      NodeAttributesManager.setAriaOrientation(
        tabListEl,
        getOrientation(tabConfig.Placement),
      );

      Array.from(tabListEl.querySelectorAll("button")).forEach((btn) => {
        if (
          btn.getAttribute("data-tab-id") === tabConfig.ActiveTab.toString()
        ) {
          NodeAttributesManager.setActiveAttributes(btn);
        }
      });

      const a11yActions = new A11YManager(tabNode, {
        enabled: true,
        activeTab: tabConfig.ActiveTab,
        bordered: tabConfig.Bordered,
        variant: tabConfig.Variants[0],
      });
      a11yActions.setup();

      const variantObserver = new MutationObserver((mutationsList) =>
        handleVariantChange(rootStyles, tabConfig, mutationsList),
      );
      variantObserver.observe(tabNode, { attributes: true });

      // Create an observer instance linked to the callback function
      const bodyBgObserver = new MutationObserver((mutationsList) =>
        handleBodyBackgroundColorChange(rootStyles, tabConfig, mutationsList),
      );
      bodyBgObserver.observe(tabNode, {
        attributes: true,
        attributeFilter: ["style", "class"],
      });

      window.addEventListener("resize", () => {
        handleWindowResize(tabKey);
      });
    }

    /* *******************************************************************
     * Main Flow
     *
     * Loop over all tabber components in the page and initialise them.
     *
     * ***************************************************************** */
    for (const key in configMap.Data) {
      const tabberConfig = configMap.Data[key];
      const containerNode = document.getElementById(` + "`" + `tabber-container-${key}` + "`" + `);

      if (!containerNode) {
        console.error(` + "`" + `Container node not found for key ${key}` + "`" + `);
        return; // Halt execution here
      }

      initializeTabberInstance(rootStyles, key, tabberConfig, containerNode);
    }
  });
}`,
		Call:       templ.SafeScript(`__templ_TabberJS_440d`, configMap),
		CallInline: templ.SafeScriptInline(`__templ_TabberJS_440d`, configMap),
	}
}

var _ = templruntime.GeneratedTemplate
