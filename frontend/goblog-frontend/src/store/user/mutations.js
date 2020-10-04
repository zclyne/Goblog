export function saveUserToken (state, userToken) {
    state.userToken = userToken
}

export function deleteUserToken (state) {
    state.userToken = ''
}
