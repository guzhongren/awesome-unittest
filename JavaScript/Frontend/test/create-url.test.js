import { createUrl } from '../src/geo_location/locateme'
describe("create-url test", () => {
  it("should return correct url given lat and lon", () => {
    const lat = 116.397;
    const lon = 39.916;

    const url = createUrl(lat, lon);

    expect(url).toBe('https://map.baidu.com/poi/@116.397,39.916');
  });

  it('should return correct url given another lat and lon', () => {
      const lat = 107.401;
      const lon = 33.425;

      const url = createUrl(lat, lon);

      expect(url).toBe('https://map.baidu.com/poi/@107.401,33.425');
  });

  it('should return empty string if lat is undefined', () => {
    const lat = undefined;
    const lon = 33.425;

    const url = createUrl(lat, lon);

    expect(url).toBe('');
  });

  it('should return empty string if lon is undefined', () => {
    const lat = 107.401;
    const lon = undefined;

    const url = createUrl(lat, lon);

    expect(url).toBe('');
  });
});
