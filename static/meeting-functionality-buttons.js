const btnStreamMicrophone = document.getElementById('btn-stream-microphone');

changerAudioStream = function() {
    let activated = false;
    return () => {
        if (activated) {
            // Synchronizer.closeAudioStream(username);
            btnStreamMicrophone.className = 'stream-functionality-button stream-functionality-microphone-off'
        } else {
            // Synchronizer.openAudioStream(username);
            btnStreamMicrophone.className = 'stream-functionality-button stream-functionality-microphone-on'
        }
        activated = !activated;
    }
}();

const btnStreamWebCamera = document.getElementById('btn-stream-web-camera');

changerVideoStream = function() {
    let activated = false;
    return () => {
        if (activated) {
            // Synchronizer.closeVideoStream(username);
            btnStreamWebCamera.className = 'stream-functionality-button stream-functionality-web-camera-off'
        } else {
            // Synchronizer.openVideoStream(username);
            btnStreamWebCamera.className = 'stream-functionality-button stream-functionality-web-camera-on'
        }
        activated = !activated;
    }
}();

btnStreamMicrophone.addEventListener('click', () => {
    changerAudioStream();
})

btnStreamWebCamera.addEventListener('click', () => {
    changerVideoStream();
})