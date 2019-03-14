import * as THREE from 'three';
import * as constants from 'constants/networkGraph';
import { getNodeCanvas, CreateTextLabelMesh } from 'utils/networkGraphUtils';

const Node = (scene, data) => {
    const node = data;

    let circle = null;
    let label = null;

    const nodeCanvas = getNodeCanvas(node);
    const nodeTexture = new THREE.Texture(nodeCanvas);
    nodeTexture.needsUpdate = true;
    const geometry = new THREE.PlaneBufferGeometry(16, 16);
    const material = new THREE.MeshBasicMaterial({
        map: nodeTexture,
        userData: {
            node,
            type: constants.NETWORK_GRAPH_TYPES.NODE
        },
        transparent: true
    });
    circle = new THREE.Mesh(geometry, material);

    label = CreateTextLabelMesh(
        node.deploymentName,
        constants.NODE_LABEL_CANVAS_SIZE,
        constants.NODE_LABEL_FONT_SIZE,
        true
    );

    scene.add(circle);
    scene.add(label);

    function setNodeOpacity(value) {
        circle.material.opacity = value;
        label.material.opacity = value;
    }

    function update() {
        const { x, y } = node;
        circle.position.set(x, y, 0);
        label.position.set(x, y - constants.NODE_LABEL_OFFSET, 0);
    }

    function getType() {
        return constants.NETWORK_GRAPH_TYPES.NODE;
    }

    function getDeploymentId() {
        return node.deploymentId;
    }

    function unhighlight() {
        setNodeOpacity(0.5);
    }

    function highlight() {
        setNodeOpacity(1);
    }

    function cleanUp() {
        scene.remove(circle);
        scene.remove(label);
        nodeTexture.dispose();
        geometry.dispose();
        material.dispose();
    }

    return {
        update,
        getType,
        getDeploymentId,
        highlight,
        unhighlight,
        cleanUp
    };
};

export default Node;
