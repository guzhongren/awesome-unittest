import { onError, onSuccess } from '../src/geo_location/locateme';

describe('onError test', () => {
    it('should set the error DOM element', () => {
        const positionError = {message: 'unable to get your location'}

				Object.defineProperty(document, 'getElementById', {
					writable: true,
					value: jest.fn().mockReturnValue({}),
				})

        onError(positionError)

				expect(document.getElementById).toHaveBeenCalled()
				expect(document.getElementById).toBeCalledTimes(1)
    });
});
