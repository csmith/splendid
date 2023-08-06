import { hmac } from '@noble/hashes/hmac';
import { sha256 } from '@noble/hashes/sha256';
import * as secp from '@noble/secp256k1';
import {get} from "svelte/store";

secp.etc.hmacSha256Sync = (k, ...m) => hmac(sha256, k, secp.etc.concatBytes(...m));

export function newPrivateKey() {
    return secp.etc.bytesToHex(secp.utils.randomPrivateKey());
}

export function getPublicKey(privateKey) {
    const privateKeyBytes = secp.etc.hexToBytes(privateKey);
    return secp.etc.bytesToHex(secp.getPublicKey(privateKeyBytes));
}

export function sign(payload, privateKey) {
    const privateKeyBytes = secp.etc.hexToBytes(privateKey);
    const hash = hmac(sha256, "splendid", new TextEncoder().encode(JSON.stringify(payload)));
    return secp.sign(hash, privateKeyBytes).toCompactHex();
}

export function createAttestation(payload, privateKey) {
    return {
        payload,
        signature: sign(payload, privateKey),
        publicKey: getPublicKey(privateKey),
    }
}

export function verify(payload, signature, publicKey) {
    const hash = hmac(sha256, "splendid", new TextEncoder().encode(JSON.stringify(payload)));
    return secp.verify(signature, hash, publicKey);
}