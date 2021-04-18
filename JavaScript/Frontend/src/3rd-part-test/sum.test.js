import jwt_decode from "jwt-decode";
import sum from "./sum";

const expectDecode = { exp: 1393286893, foo: "bar", iat: 1393268893 };

jest.mock("jwt-decode", () => {
	return jest.fn().mockImplementation(() => {
		return "test";
	});
});

test("adds 1+2 to equal 3", () => {
	expect(sum(1, 2)).toBe(3);
});

it("jwt", () => {
	var token =
		"eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJmb28iOiJiYXIiLCJleHAiOjEzOTMyODY4OTMsImlhdCI6MTM5MzI2ODg5M30.4-iaDojEVl0pJQMjrbM1EzUIfAZgsbK_kgnVyVxFSVo";

	const decoded = jwt_decode(token);
	expect(decoded).toBe("test");
});
