package webrtc

import (
    "github.com/pion/webrtc/v3"
    "log"
)

func HandleOffer(offer webrtc.SessionDescription) (*webrtc.SessionDescription, error) {
    peerConnection, err := CreatePeerConnection()
    if err != nil {
        return nil, err
    }

    err = peerConnection.SetRemoteDescription(offer)
    if err != nil {
        return nil, err
    }

    answer, err := peerConnection.CreateAnswer(nil)
    if err != nil {
        return nil, err
    }

    err = peerConnection.SetLocalDescription(answer)
    if err != nil {
        return nil, err
    }

    return &answer, nil
}