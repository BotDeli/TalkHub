const btnStreamMicrophone = document.getElementById('btn-stream-microphone');
const btnStreamWebCamera = document.getElementById('btn-stream-web-camera');

class VideoAudioController {
    constructor(video, audio) {
        this.videoActive = video;
        this.audioActive = audio;
        if (this.videoActive) {
            btnStreamWebCamera.className = 'stream-functionality-button stream-functionality-web-camera-on'
        } else {
            btnStreamWebCamera.className = 'stream-functionality-button stream-functionality-web-camera-off'
        }

        if (this.audioActive) {
            btnStreamMicrophone.className = 'stream-functionality-button stream-functionality-microphone-on'
        } else {
            btnStreamMicrophone.className = 'stream-functionality-button stream-functionality-microphone-off'
        }

        btnStreamWebCamera.addEventListener('click', () => {
            if (this.videoActive) {
                this.offVideo();
            } else {
                this.onVideo();
            }
        })
        btnStreamMicrophone.addEventListener('click', () => {
            if (this.audioActive) {
                this.offAudio();
            } else {
                this.onAudio();
            }
        })
    }

    onVideo() {
        if (!this.videoActive) {
            this.videoActive = true;
            updateTrackTransfer('video', true)
            btnStreamWebCamera.className = 'stream-functionality-button stream-functionality-web-camera-on'
        }
    }

    offVideo() {
        if (this.videoActive) {
            this.videoActive = false;
            updateTrackTransfer('video', false)
            btnStreamWebCamera.className = 'stream-functionality-button stream-functionality-web-camera-off'
        }
    }

    onAudio() {
        if (!this.audioActive) {
            this.audioActive = true;
            updateTrackTransfer('audio', true)
            btnStreamMicrophone.className = 'stream-functionality-button stream-functionality-microphone-on'
        }
    }

    offAudio() {
        if (this.audioActive) {
            this.audioActive = false;
            updateTrackTransfer('audio', false)
            btnStreamMicrophone.className = 'stream-functionality-button stream-functionality-microphone-off'
        }
    }
}