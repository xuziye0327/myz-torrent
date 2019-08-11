window.onload = function () {
    fetchFiles()
    fetchStates()
    setInterval(fetchStates, 5 * 1000);
};

function fetchStates() {
    fetch("/download", {
        method: "GET",
    }).then(resp => {
        if (resp.status !== 200) {
            alert("get index info err!");
            return null;
        }

        return resp.json()
    }).then(updateStates)
}

function fetchFiles() {
    fetch("/file", {
        method: "GET",
    }).then(resp => {
        if (resp.status !== 200) {
            alert("get index info err!");
            return null;
        }

        return resp.json()
    }).then(updateFiles)
}

function updateStates(states) {
    const body = document.querySelector("#downloads")
    while (body.firstChild) {
        body.removeChild(body.firstChild)
    }

    const template = document.querySelector("#downloads_template")
    for (s of states) {
        const clone = template.content.cloneNode(true)

        const name = clone.querySelector("p")
        name.textContent = s.name

        const opts = clone.querySelectorAll("button");
        for (opt of opts) {
            opt.value = s.id
        }

        const progress = clone.querySelector("#download_progress")
        progress.style.width = s.state.percent + "%"
        progress.textContent = s.state.percent.toFixed(2) + "%"

        body.appendChild(clone)
    }
}

function download() {
    const url = document.getElementById("url");
    const val = url.value.trim();
    if (val.length === 0) {
        return;
    }

    fetch("/download", {
        method: 'POST',
        body: JSON.stringify([val]),
    }).then(resp => {
        if (resp.status === 200) {
            alert("success!")
            return
        }
        resp.text().then(alert)
    })
}

function updateFiles(files) {
    if (files.length === 0) {
        return
    }

    const body = document.querySelector("#files")
    while (body.firstChild) {
        body.removeChild(body.firstChild)
    }

    const template = document.querySelector("#files_template")
    const clone = template.content.cloneNode(true)
    const ul = clone.querySelector("ul")

    const lis = updateChildFiles("", files)
    for (li of lis) {
        ul.appendChild(li)
    }

    body.appendChild(clone)
}

function updateChildFiles(prefix, files) {
    let id = 0
    const ret = []
    for (f of files) {
        const curId = prefix + '-' + (id++)

        const tmp = document.querySelector("#files_li_template")
        const t = tmp.content.cloneNode(true)
        const file_link = t.querySelector("#file_link")
        file_link.href = "/file/" + pathEncode(f.full_path)
        file_link.textContent = f.name

        if (f.is_dir === true) {
            const show_file_child = t.querySelector("#show_file_child")
            show_file_child.id = "show_file_child" + curId
            show_file_child.textContent = "+"
            show_file_child.setAttribute("onclick", "showFileChild('" + curId + "')")

            const ul = t.querySelector("#file_child")
            ul.id = "file_child" + curId
            const childs = updateFiles(curId, f.childs)
            for (c of childs) {
                ul.appendChild(c)
            }
        }

        ret.push(t)
    }
    return ret
}

function showFileChild(id) {
    const file_child = document.getElementById("file_child" + id)
    const show_file_child = document.getElementById("show_file_child" + id)
    console.log(file_child.style.display)
    if (file_child.style.display === "none") {
        file_child.style.display = ""
        show_file_child.textContent = "-"
    } else {
        file_child.style.display = "none"
        show_file_child.textContent = "+"
    }
}

// firstly convert to urlEncode
// then using base64 encode
function pathEncode(path) {
    path = encodeURI(path);
    path = window.btoa(path);
    return path;
}