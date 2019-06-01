function downloadMagnet() {
    const form = document.getElementById("magnet-form");
    const formInputDivs = form.getElementsByTagName("div");

    const magnets = [];
    for (let i = 0; i < formInputDivs.length; ++i) {
        const div = formInputDivs.item(i);
        const input = div.getElementsByTagName("input").item(0);

        const val = input.value.trim();
        if (val.length !== 0) {
            magnets.push(val);
        }
    }

    if (magnets.length === 0) {
        console.error("magnets are empty!");
        return false;
    }

    return fetch("/torrent/magnet", {
        method: 'POST',
        body: JSON.stringify(magnets),
    }).then(resp => {
        if (resp.status === 200) {
            alert("success!");
            refreshMagnetLine();
            return true;
        } else {
            return resp.text()
                .then(msg => {
                    alert(msg);
                    return false;
                });
        }
    });
}

function refreshMagnetLine() {
    const form = document.getElementById("magnet-form");
    form.innerHTML = "";
    addNewMagnetLine();
}

function addNewMagnetLine() {
    const form = document.getElementById("magnet-form");
    const formInputDivs = form.getElementsByTagName("div");
    const newDivs = [];

    for (let i = 0; i < formInputDivs.length; ++i) {
        const div = formInputDivs.item(i);
        const input = div.getElementsByTagName("input").item(0);

        newDivs.push(generateMagnetLine(i, input.value, false, false));
    }

    newDivs.push(generateMagnetLine(formInputDivs.length, "", true, formInputDivs.length === 0));

    form.innerHTML = newDivs.join("");
}

function deleteNewMagnetLine(id) {
    const form = document.getElementById("magnet-form");
    const formInputDivs = form.getElementsByTagName("div");

    const filterDivs = [];
    for (let i = 0; i < formInputDivs.length; ++i) {
        const div = formInputDivs.item(i);

        if (div.id === id) {
            continue;
        }

        filterDivs.push(div);
    }

    const newDivs = [];
    for (let i = 0; i < filterDivs.length; ++i) {
        const div = filterDivs[i];
        const input = div.getElementsByTagName("input").item(0);

        newDivs.push(generateMagnetLine(i, input.value, i === filterDivs.length - 1, filterDivs.length === 1))
    }

    form.innerHTML = newDivs.join("");
}

function generateMagnetLine(index, value, isLast, isOnly) {
    let ret = '<div class="form-inline mr-sm-2" id="magnet-form-input-' + index + '">';
    ret += '<input type="text" class="form-control col mr-sm-2" id="magnet-' + index + '" placeholder="magnet:?xt=urn:btih:" value="' + value + '">';
    if (isLast) {
        ret += '<button type="button" class="btn btn-primary mr-sm-2" onclick="addNewMagnetLine()">+</button>';
    }
    if (!isOnly) {
        ret += '<button type="button" class="btn btn-danger mr-sm-2" onclick="deleteNewMagnetLine(\'magnet-form-input-' + index + '\')">-</button>';
    }
    ret += '</div>';
    return ret;
}

window.onload = function () {
    addNewMagnetLine();

    addSubFiles("files", "");
};

function addSubFiles(divId, path) {
    path = path.trim();
    if (path.length !== 0) {
        path = pathEncode(path);
    }

    fetch("/file?path=" + path, {
        method: "GET",
    }).then(resp => {
        if (resp.status !== 200) {
            alert("get files err!");
            return null;
        }

        return resp.json()
    }).then(res => {
        if (res === null) {
            return;
        }

        const curDiv = document.getElementById(divId);

        let index = 0;
        let html = '<ul class="list-group">';
        for (let f of res) {
            let line = '<li class="list-group-item">';
            line += '<div class="form-inline">';
            line += '<div class="col mr-sm-2"><a href="/file/' + pathEncode(f.full_path) + '">' + f.name + '</a></div>';

            let subId = curDiv.id + '-' + index;
            if (f.is_dir === true) {
                line += '<a id="'+ subId + '-option' +'" href="javascript:void(0);" onclick="addSubFiles(\'' + subId + '\', \'' + f.full_path + '\')">+</a>';
            }

            line += '</div>';

            if (f.is_dir === true) {
                line += '<div class="list-group" id="' + subId + '"></div>';
                index += 1;
            }

            html += line;
        }
        html += '</ul>';

        curDiv.innerHTML = html;

        // update current div option
        const curOption = document.getElementById(divId + "-option");
        if (curOption === null) {
            return;
        }

        let curOnclick = curOption.getAttribute("onclick");
        curOnclick = curOnclick.replace("addSubFiles", "removeSubFiles");

        curOption.setAttribute("onclick", curOnclick);
        curOption.innerText = '-';
    });
}

function removeSubFiles(divId) {
    document.getElementById(divId).innerHTML = "";

    const curOption = document.getElementById(divId + "-option");

    let curOnclick = curOption.getAttribute("onclick");
    curOnclick = curOnclick.replace("removeSubFiles", "addSubFiles");

    curOption.setAttribute("onclick", curOnclick);
    curOption.innerText = '+';
}

// firstly convert to urlEncode
// then using base64 encode
function pathEncode(path) {
    path = encodeURI(path);
    path = window.btoa(path);
    return path;
}