import { onError, onSuccess } from '../src/geo_location/locateme';

describe('onError test', () => {
    it('should set the error DOM element', () => {
        const positionError = {message: 'unable to get your location'}

        onError(positionError);

        expect(document.getElementById('error').innerHTML).toEqual(positionError.message);

    });
});