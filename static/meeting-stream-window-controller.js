const aspectRation = 4/3;

const userList = document.getElementById('users-list');
const streamUsersList = document.getElementById('stream-users-list');
class StreamWindowController {
    constructor() {
        this.streams = new Map();
        this.videos = new Map();
        this.users = new Map();
        this.count = 0;

        window.addEventListener('resize',() => {
            this.optimizeResizeWindows();
        });
    }

    startNewStream(userID, username, stream) {
        if (userID === undefined || username === undefined || stream === undefined) {
            return;
        }

        const savedVideo = this.videos.get(userID);

        if (savedVideo !== undefined && savedVideo !== null) {
            this.streams.set(userID, stream);
            savedVideo.srcObject = stream;
            return;
        }

        let video = document.createElement('video');
        video.className = 'stream-user';
        video.style.width = '400px';
        video.style.height = '300px';
        video.autoplay = true;
        video.srcObject = stream;
        streamUsersList.appendChild(video);

        let user = document.createElement('div');
        user.className = 'users-list-user';

        let div = document.createElement('div');
        div.className = 'profile-image';
        user.appendChild(div);

        div = document.createElement('div');
        div.className = 'profile-username';
        div.innerText = username;
        user.appendChild(div);

        userList.appendChild(user);

        this.streams.set(userID, stream);
        this.videos.set(userID, video);
        this.users.set(userID, user);
        this.count++;

        this.optimizeResizeWindows();
    }

    endStream(userID) {
        const savedVideo = this.videos.get(userID);
        if (savedVideo === undefined || savedVideo === null) {
            return;
        }
        savedVideo.remove();
        const savedUser = this.users.get(userID);
        savedUser.remove();

        this.streams.delete(userID);
        this.videos.delete(userID);
        this.users.delete(userID);
        this.count--;

        this.optimizeResizeWindows();
    }

    optimizeResizeWindows() {
        let freeWidth = Math.floor(streamUsersList.clientWidth) - 160;
        let freeHeight = Math.floor(streamUsersList.clientHeight) - 10;
        let width;
        let height;
        if (this.count === 1) {
            [width, height] = findAspectRatioSize(freeWidth, freeHeight);
        } else if (this.count === 2) {
            let [w1, h1] = findAspectRatioSize(freeWidth/2, freeHeight);
            let [w2, h2] = findAspectRatioSize(freeWidth, freeHeight/2);
            width = max(w1, w2);
            height = max(h1, h2);
        } else {
            let [w1, h1] = findAspectRatioSize(freeWidth/4, freeHeight);
            let [w2, h2] = findAspectRatioSize(freeWidth, freeHeight/4);
            let [w3, h3] = findAspectRatioSize(freeWidth/2, freeHeight/2);
            width = max(w1, w2);
            height = max(h1, h2);
            width = max(width, w3);
            height = max(height, h3);
        }

        this.videos.forEach((key, value) => {
            key.style.width = width + 'px';
            key.style.height = height + 'px';
        })
    }
}

function findAspectRatioSize(width, height) {
    if (width <= height) {
        height = width / aspectRation;
    } else {
        width = height * aspectRation;
    }

    return [Math.floor(width), Math.floor(height)];
}

function max(n1, n2) {
    if (n1 > n2) {
        return n1;
    }
    return n2
}