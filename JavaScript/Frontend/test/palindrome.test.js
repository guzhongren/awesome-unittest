import { isPalindrome } from '../src/palindrome';
describe('palindrome test', () => {
    it('should pass this canary test', () => {
        expect(true).toBe(true);
    });

    it('should return true for argument mom', () => {
        expect(isPalindrome('mom')).toBe(true);
    });

    it('should return false for argument momu', () => {
        expect(isPalindrome('momu')).toBe(false);
    });

    it('should return false when argument is an empty string', () => {
        expect(isPalindrome('')).toBe(false);
    });

    it('should return false for argument with only two spaces', () => {
        expect(isPalindrome('  ')).toBe(false);
    });

    it('should throw an exception if argument is missing', () => {
        const call = () => {
            isPalindrome();
        };

        expect(call).toThrow('Invalid argument');
    });

    it('should throw an exception if argument is a', () => {
        const call = () => {
            isPalindrome('a');
        };
        
        expect(call).toThrow('Single word');
    });
});