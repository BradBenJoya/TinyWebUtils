const element = document.getElementById("file");

element.addEventListener("cancel", () => {
    console.log("Cancelled file upload");
});

element.addEventListener("change", async () => {
    if (element.files.length === 0) return;

    const file = element.files[0];
    const formData = new FormData();
    formData.append("files", file);

    try {
        const res = await fetch("/corruptor/api", {
            method: "POST",
            body: formData,
        });
        if (!res.ok) throw new Error("Request failed: " + res.status);

        const blob = await res.blob();
        const url = URL.createObjectURL(blob);
        const a = document.createElement("a");
        a.href = url;
        a.download = file.name;
        document.body.appendChild(a);
        a.click();
        a.remove();
        URL.revokeObjectURL(url);
    } catch (err) {
        console.error(err);
    }
});