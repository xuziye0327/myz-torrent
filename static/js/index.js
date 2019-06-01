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
    let ret = '<div class="form-inline my-1" id="magnet-form-input-' + index + '">';
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
};