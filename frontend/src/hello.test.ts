it("Hello vitest", () => {
  const hello = "Hello vitest"
  expect(hello).toBe("Hello vitest")
})

it("Not jest", () => {
  const hello = "Hello vitest"
  expect(hello).not.toBe("Hello jest")
})
