import jwt_decode from "jwt-decode";
jest.mock("jwt-decode", () => {
	return jest.fn().mockImplementation(() => {
		return "test";
	});
});

it("jwt", () => {
	// const expectDecode = { exp: 1393286893, foo: "bar", iat: 1393268893 };
	var token =
		"eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJmb28iOiJiYXIiLCJleHAiOjEzOTMyODY4OTMsImlhdCI6MTM5MzI2ODg5M30.4-iaDojEVl0pJQMjrbM1EzUIfAZgsbK_kgnVyVxFSVo";

	const decoded = jwt_decode(token);
	expect(decoded).toBe("test");
});
