const aspectRation = 4/3;

class StreamSynchronizer {
    constructor() {
        this.VideOutput = {};
        this.VideoStreamActive = {};
        this.AudioStreamActive = {};
        this.countStreams = 0;
    }

    addVideoOutput(userID, videOutput) {
        this.VideOutput[userID] = videOutput;
        this.VideoStreamActive[userID] = null;
        this.AudioStreamActive[userID] = null;
        this.countStreams++;
    }

    synchronize() {
        let width = streamUsersList.clientWidth;
        let height = streamUsersList.clientHeight;

        if (width >= height) {
            width /= aspectRation;
        } else {
            height *= aspectRation;
        }

        if (this.countStreams >= 2) {
            if (this.countStreams >= 5) {
                if (this.countStreams >= 10) {
                    width /= 4;
                } else {
                    width /= 3;
                }
            } else {
                width /= 2;
            }
        }

        if (this.countStreams >= 3) {
            if (this.countStreams >= 7) {
                height /= 3;
            } else {
                height /= 2;
            }
        }

        Object.values(this.VideOutput).forEach((output) => {
            output.style.maxWidth = width + 15 + 'px';
            output.style.maxHeight = height - 5 + 'px';
        })
    }

    openVideoStream(userID) {
        if (this.VideoStreamActive[userID] === null) {
            navigator.mediaDevices.getUserMedia({video: true}).
                then(stream => {
                    this.VideoStreamActive[userID] = stream;
                    this.VideOutput[userID].srcObject = stream;
                });
        }
    }

    openAudioStream(userID) {
        if (this.AudioStreamActive[userID] === null) {
            navigator.mediaDevices.getUserMedia({audio: true}).
            then(stream => {
                this.AudioStreamActive[userID] = stream;
                this.VideOutput[userID].srcObject = stream;
            });
        }
    }

    closeVideoStream(userID) {
        if (this.VideoStreamActive[userID] !== null) {
            this.VideoStreamActive[userID].getVideoTracks()[0].stop();
            this.VideoStreamActive[userID] = null;
        }
    }

    closeAudioStream(userID) {
        if (this.AudioStreamActive[userID] !== null) {
            this.AudioStreamActive[userID].getAudioTracks()[0].stop();
            this.AudioStreamActive[userID] = null;
        }
    }
}