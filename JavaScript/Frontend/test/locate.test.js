import { locate } from '../src/geo_location/locateme';
describe('locate test', () => {
    it('should register handlers with getCurrentPosition', () => {
        let originNavigator = navigator.geolocation;

        navigator.geolocation = {
            getCurrentPosition: jest.fn()
        };

        locate();
        expect(navigator.geolocation.getCurrentPosition).toHaveBeenCalled()

        navigator.geolocation = originNavigator;
    });
});