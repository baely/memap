class MapView {
    constructor(root) {
        this.root = root;
    }

    init() {
        this.initCanvas();
        this.initInputs();
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
            engine.buttonPress("Search");
        });

        searchPanel.appendChild(searchInput);
        searchPanel.appendChild(searchButton);

        this.root.appendChild(searchPanel);
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
