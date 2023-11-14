const eyes = document.getElementsByClassName('eye');
for (let i = 0; i < eyes.length; i++) {
    eyes[i].addEventListener('click', (e) => {
        let nodeFor = eyes[i].getAttributeNode('for');
        if (nodeFor !== null) {
            let node = document.getElementById(nodeFor.value);
            if (node.getAttributeNode('type').value === 'password') {
                node.setAttribute('type', 'text');
                eyes[i].className = 'eye';
            } else {
                node.setAttribute('type', 'password');
                eyes[i].className = 'eye closed-eye';
            }
            e.stopPropagation();
        }
    })
}