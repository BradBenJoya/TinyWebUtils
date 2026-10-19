const element = document.createElement("input")
element.type = "file"

element.addEventListener("cancel", () => {
    console.log("Canceld file upload")
})

element.addEventListener("change", () => {
    if (element.files.length === 1) {
        console.log("File selected: ", element.files[0])
    }
})

element.click()