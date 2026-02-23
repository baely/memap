class MapView {
    constructor(root) {
        this.root = root;
    }

    init() {
        this.initSidebarInputs();
        this.initInfoPanel();
        this.initCanvas();
        this.initInputs();
    }

    initCanvas() {
        const canvas = document.createElement("canvas");
        this.root.appendChild(canvas);

        this.ctx = canvas.getContext("2d");
        engine.initRenderer(canvas, this);

        const ratio = window.devicePixelRatio || 1;
        let refreshSize = () => {
            canvas.width = this.root.clientWidth * ratio;
            canvas.height = this.root.clientHeight * ratio;

            canvas.style.width = this.root.clientWidth + "px";
            canvas.style.height = this.root.clientHeight + "px";

            this.ctx.scale(ratio, ratio);

            // engine.updateViewport(canvas.width, canvas.height, ratio);
            engine.updateViewport(this.root.clientWidth, this.root.clientHeight);
            engine.draw();
        }

        window.addEventListener("resize", refreshSize);
        refreshSize();
    }

    initInputs() {
        const searchPanel = document.createElement("div");
        searchPanel.classList.add("search")

        const searchInput = document.createElement("input");
        const searchButton = document.createElement("button");

        searchInput.id = "search-input"
        searchInput.placeholder = "Search Bailey Maps";
        searchInput.type = "text";
        searchButton.textContent = "🔎";

        searchButton.addEventListener("click", () => {
            engine.buttonPress("search");
        });

        searchPanel.appendChild(searchInput);
        searchPanel.appendChild(searchButton);

        this.root.appendChild(searchPanel);
    }

    initSidebarInputs() {
        const sidebarPanel = document.createElement("div");
        sidebarPanel.id = "menu-panel";
        sidebarPanel.classList.add("menu-panel");

        this.root.appendChild(sidebarPanel);
    }

    initInfoPanel() {
        const infoPanel = document.createElement("div");
        infoPanel.id = "info-panel";
        infoPanel.classList.add("info-panel");

        const titleText = document.createElement("div");
        titleText.classList.add("title");

        const linkText = document.createElement("div");
        linkText.classList.add("link");

        const linkDOM = document.createElement("a");
        linkDOM.setAttribute("target", "_blank");
        linkText.appendChild(linkDOM);

        const descriptionText = document.createElement("div");
        descriptionText.classList.add("description");

        infoPanel.appendChild(titleText);
        infoPanel.appendChild(linkText);
        infoPanel.appendChild(descriptionText);

        let outerCallback;
        const doCallback = (e) => {
            if (e.key === "Enter") {
                e.target.innerText = e.target.innerText.slice(0, -1);
                e.target.blur();
            }
            outerCallback(
                titleText.innerText,
                linkDOM.innerText,
                descriptionText.innerText,
            );
        }

        this.showInfoPanel = (title, link, description) => {
            titleText.innerText = title;
            linkDOM.innerText = link;
            linkDOM.setAttribute("href", link);
            descriptionText.innerText = description;

            titleText.contentEditable = "false";
            linkText.contentEditable = "false";
            descriptionText.contentEditable = "false";

            linkText.style.display = link == null ? "none" : "block";
            descriptionText.style.display = description == null ? "none" : "block";

            infoPanel.style.display = "block";
        }
        this.showEditPanel = (title, link, description, callback) => {
            titleText.innerText = title;
            linkDOM.innerText = link;
            linkDOM.removeAttribute("href");
            descriptionText.innerText = description;

            titleText.contentEditable = "plaintext-only";
            linkDOM.contentEditable = "plaintext-only";
            descriptionText.contentEditable = "plaintext-only";

            outerCallback = callback;
            titleText.addEventListener("keyup", doCallback);
            if (link == null) {
                linkText.style.display = "none";
            } else {
                linkText.style.display = "none";
                linkDOM.addEventListener("keyup", doCallback);
            }
            if (description == null) {
                descriptionText.style.display = "none";
            } else {
                descriptionText.style.display = "block";
                descriptionText.addEventListener("keyup", doCallback);
            }

            infoPanel.style.display = "block";
        }
        this.hideInfoPanel = () => {
            infoPanel.style.display = "none";

            titleText.removeEventListener("keyup", doCallback);
            linkText.removeEventListener("keyup", doCallback);
            descriptionText.removeEventListener("keyup", doCallback);
            outerCallback = false;
        }

        this.hideInfoPanel();
        this.root.appendChild(infoPanel);
    }

    drawBatch(rawBatch) {
        const batch = JSON.parse(rawBatch);

        for (let item of batch) {
            switch (item[0]) {
                case "set":
                    this.ctx[item[1]] = item[2];
                    break;
                case "call":
                    this.ctx[item[1]](...item.slice(2));
                    break;
            }
        }
    }
}

const root = document.querySelector('#root');

const map = new MapView(root);

const go = new Go();
WebAssembly.instantiateStreaming(fetch("./maps.wasm"), go.importObject)
    .then((result) => {
        go.run(result.instance);
        map.init();
    });
