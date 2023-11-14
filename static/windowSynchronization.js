class Synchronization {
    constructor(hideElement, defaultDisplay, targetElement, defaultDVW, limitWidth) {
        this.hideElement = hideElement;
        this.defaultDisplay = defaultDisplay;
        this.targetElement = targetElement;
        this.defaultDVW = defaultDVW;
        this.limitWidth = limitWidth;
    }

    synchronize() {
        if (window.innerWidth <= this.limitWidth) {
            this.hideElement.style.display = 'none';
            this.targetElement.style.width = '100dvw';
        } else {
            this.hideElement.style.display = this.defaultDisplay;
            this.targetElement.style.width = this.defaultDVW;
        }
    }
}

const informationPanel = document.getElementById('information-panel');
const defaultDisplay = informationPanel.style.display;
const targetElement = document.getElementById('registration-form');
const defaultDVW = targetElement.style.width;
const limit = 1000;

const sync = new Synchronization(informationPanel, defaultDisplay, targetElement, defaultDVW, limit);
sync.synchronize();

window.addEventListener('resize', () => {sync.synchronize()});