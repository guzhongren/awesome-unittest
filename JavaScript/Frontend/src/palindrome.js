export const isPalindrome = (phrase) => {
    if (phrase === undefined) {
        throw new Error('Invalid argument')
    }

    if (phrase.length === 1) {
        throw new Error('Single word')
    }

    return phrase.trim().length > 0 && phrase.split('').reverse().join('') === phrase;
}