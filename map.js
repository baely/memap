class MapView {
    constructor(root) {
        this.root = root;
    }

    init(editMode) {
        this.editMode = editMode;
        this.initCanvas();
        this.initInputs();

        if (editMode) {
            this.initEditInputs();
        } else {
            this.initSidebarInputs();
        }
    }

    initCanvas() {
        const canvas = document.createElement("canvas");
        this.root.appendChild(canvas);

        this.ctx = canvas.getContext("2d");
        engine.initRenderer(canvas, this.drawBatch.bind(this), this.ctx);

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
        sidebarPanel.classList.add("edit");

        const editButton = document.createElement("button");
        editButton.textContent = "✏️";
        editButton.addEventListener("click", () => {
            window.location.search = "?edit";
        });
        sidebarPanel.appendChild(editButton);

        this.root.appendChild(sidebarPanel);
    }

    initEditInputs() {
        const editPanel = document.createElement("div");
        editPanel.classList.add("edit");

        const downloadButton = document.createElement("button");
        downloadButton.textContent = "⬇️";
        downloadButton.addEventListener("click",  () => {
            engine.buttonPress("download");
        });
        editPanel.appendChild(downloadButton);

        const closeButton = document.createElement("button");
        closeButton.textContent = "❌";
        closeButton.addEventListener("click", () => {
            window.location.search = "";
        });
        editPanel.appendChild(closeButton);

        this.root.appendChild(editPanel);
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
const params = new URLSearchParams(window.location.search);
const editMode = params.has("edit");
const wasmFile = editMode ? "./editor.wasm" : "./maps.wasm";

if (editMode) {
    document.title = "Bailey Maps - Editor";
}

const map = new MapView(root);

const go = new Go();
WebAssembly.instantiateStreaming(fetch(wasmFile), go.importObject)
    .then((result) => {
        go.run(result.instance);
        map.init(editMode);
    });
