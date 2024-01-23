wit_bindgen::generate!({
    world: "host",
    exports: {
        world: Decider,
    },
});

struct Decider;

impl Guest for Decider {
  fn run() {
    print("Hello, world!");
  }
}
