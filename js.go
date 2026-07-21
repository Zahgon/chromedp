package chromedp

import (
	_ "embed"
)

var (

	//go:embed js/text.js
	textJS string

	//go:embed js/textContent.js
	textContentJS string

	//go:embed js/blur.js
	blurJS string

	//go:embed js/submit.js
	submitJS string

	//go:embed js/reset.js
	resetJS string

	//go:embed js/attribute.js
	attributeJS string

	//go:embed js/setAttribute.js
	setAttributeJS string

	//go:embed js/visible.js
	visibleJS string

	//go:embed js/getClientRect.js
	getClientRectJS string

	//go:embed js/waitForPredicatePageFunction.js
	waitForPredicatePageFunction string
)
