class HideSwitcher {
    constructor(object, defaultDisplay) {
        this.object = object;
        if (typeof defaultDisplay === 'undefined') {
            this.display = 'unset';
        } else {
            this.display = defaultDisplay;
        }
        this.active = false;
    }

    show() {
        if (!this.active) {
            this.object.style.display = this.display;
            this.active = true;
        }
    }

    hide() {
        if (this.active) {
            this.object.style.display = "none";
            this.active = false;
        }
    }

    replace() {
        if (this.active) {
            this.hide();
        } else {
            this.show();
        }
    }
}