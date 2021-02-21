import { setLocation } from '../src/geo_location/locateme'
describe('setLocation test', () => {
    it('should set the URL into location of window', () => {
        const windowStub = {};
        const url = 'https://map.baidu.com/poi/@107.401,33.425';
        
        setLocation(windowStub, url);

        expect(windowStub.location).toBe(url);
    });
});