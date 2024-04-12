# MapXR
A custom mapping tool for the desktop, extending the capabilities of the TapXR's built-in [map.tapwithus.com](map.tapwithus.com) mappings, enabling support for many layers, a wide array of custom actions, and two handed integrations.

### How It Works
When you run the executable, MapXR will connect to any paired Tap devices it can find, set them to `controller` mode which will ignore any mappings you already have on the device and it will instead interpet your taps according to the mappings specified in the `config.toml` file.

The downside to this approach is that the map is specific to the desktop you are connecting to since it requires MapXR to be running, and the map doesn't follow you around from device to device like the built in tap map does.

However, the upside is that it enables custom behavior you cannot achieve otherwise. It also makes changing and tweaking your map as simple as editing a file.

### Intent
The TapXR is such a niche product, and needing additional features not provided in the default tap map tools, is even more niche. Owning two TapXR's and wanting them to work together is even yet more niche. All this to say, I'm making this for myself, but providing it and the source for free under an MIT license just in case it is useful to the 3 other people in the world who might want such a thing. Feel free to report bugs and issues as it's good to record them, but I make no guarantee to address them if they don't impact me. Feel free to fork and do with it whatever you want.

I do fully intend to use MapXR as a way to utilize the TapXr as a real keyboard replacement and productivity tool, and not just a fun gimmick. This means it should enable you to access every key normally available to a full sized keyboard in as few taps as possible. In theory, with 2 TapXR's working in tandem, this means you would have (2^10)-1, aka 1023 unique single tap combinations to map key presses to. Realistically though, that's too many combinations to memorize and will likely require an intelligently designed system of layering and simple two handed combinations to be usable.

I have basically no plans to create a UI for a variety of reasons but I will make efforts to make the `config.toml` as simple to edit as possible. I am taking efforts to allow MapXR to run on Linux, Windows and Mac, though I have to Mac to test with. Cross compatibility is the reason Go was chosen as the language to build this in, but this is the first serious project I've written in go so take any code written with that in mind. 

While the SDK hasn't changed between the Tap Strap and the TapXr, and the Tap Strap should work, I'm not intentionally testing or working on ensuring that it does work with the Strap.


## Features
- [ ] \*Infinite number of Layers [^1]
- [ ] Available Bindable Actions:
    - [x] Key Press
    - [ ] Multiple Key Presses
    - [ ] Hold Key Press
        - [ ] Until tapped again
        - [ ] For specified duration
    - [ ] Hold Multiple Key Presses
        - [ ] Until tapped again
        - [ ] For specified duration
    - [ ] Modifier Key Press (auto releases after next key press)
    - [ ] Layer Shifts
    - [ ] Temporary Layer Shifts (auto switches back)
    - [ ] Timed Layer Shifts (shifts back after X milliseconds)
    - [ ] Macro (sequence of key presses and delays)
    - [ ] Command (a raw command/process you want your OS to perform)
    - [ ] Toggle Standy - pauses all input / tap handling but maintains connection. Good for when you want to drink your coffee or do something with your hands without sending errant taps and you don't want to disconnect just for that.
- [ ] Support Single, Double, Triple, and Quadruple Tap Bindings
- [ ] Support for defining "Combos", a sequence of taps done within a time window that will trigger an action
- [ ] Support For Two Taps
- [ ] Auto switch between one/two handed map based on # of Taps connected
- [ ] Two handed map can either just mirror single hand map or use custom two handed defined map
- [ ] Configurable Multi-tap handling [^2]



[^1]: \* not actually, just larger than you'd ever need

[^2]: Normally, when the TapXR needs to handle a double/triple tap, it will immediately type the character, then when the second tap comes in and it recognizes it as a double tap, it deletes the chacter and then sends the double tap character. This causes a problem for any key combination that use double/triple taps because sending and then deleting interupts the combination. I will have a toggle to switch from this delete based behavior to a delay based behavior where, it will simply either wait to send the first character until a timeout or until a second tap is sent that doesn't match. This will cause small delays in some characters appearing on screen which may be annoying, but it's the only way to ensure that only the desired key is sent, preserving combinations. Ideally, though, with infinite layers, you should be able to keep everything on single taps and avoid multi-taps entirely. This is even more true if you plan to use 2 taps.

## Building MapXR
Linux requires gcc to do a go build

## Config
details explaining how to edit the config coming...