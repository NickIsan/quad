const generateButton = document.getElementById("generate")
const widthInput = document.getElementById("width")
const heightInput = document.getElementById("height")
const output = document.getElementById("output")
const quadType = document.getElementById("quadType")

generateButton.addEventListener("click", async function () {
	const x = widthInput.value
	const y = heightInput.value
	const selectedQuad = quadType.value

	const response = await fetch(
		"/api/generate?x=" + x + "&y=" + y + "&quad=" + selectedQuad
	)

	const result = await response.text()

	let styledResult = ""

	for (const char of result) {
		if (char === "o") {
			styledResult += '<span class="corner">o</span>'
		} else if (char === "-") {
			styledResult += '<span class="horizontal">-</span>'
		} else if (char === "|") {
			styledResult += '<span class="vertical">|</span>'
		} else if (char === "A") {
			styledResult += '<span class="letter-a">A</span>'
		} else if (char === "B") {
			styledResult += '<span class="letter-b">B</span>'
		} else if (char === "C") {
			styledResult += '<span class="letter-c">C</span>'
		} else if (char === "/") {
			styledResult += '<span class="slash">/</span>'
		} else if (char === "\\") {
			styledResult += '<span class="slash">\\</span>'
		} else if (char === "*") {
			styledResult += '<span class="star">*</span>'
		} else {
			styledResult += char
		}
	}

	output.innerHTML = styledResult
})
