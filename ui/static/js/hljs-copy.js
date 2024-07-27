document.addEventListener("DOMContentLoaded", (event) => {
  const buttons = document.getElementsByClassName("hljs-copy")
  for (const button of buttons) {
    button.addEventListener("click", function () {
      this.innerText = "Copying.."
      const code = this.parentElement.getElementsByTagName("code")[0].innerText
      navigator.clipboard.writeText(code).then(() => {
        this.innerText = "Copied!"
        setTimeout(() => {
          this.innerText = "Copy"
        }, 1000)
      })
    })
  }
})
