export namespace audiodevice {
	
	export class AudioDevice {
	    ID: string;
	    Name: string;
	    Selected: boolean;
	
	    static createFrom(source: any = {}) {
	        return new AudioDevice(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.ID = source["ID"];
	        this.Name = source["Name"];
	        this.Selected = source["Selected"];
	    }
	}

}

export namespace keycodes {
	
	export enum KeycodeGroup {
	    ALL = 0,
	    CHARACTERS = 1,
	    NUMBERS = 2,
	    NUMPAD = 3,
	    MODIFIERS = 4,
	    ARROWS = 5,
	    SPECIAL = 6,
	    FUNCTIONKEYS = 7,
	}
	export class Keycode {
	    Rawcode: number;
	    Display: string;
	    JSCode: string;
	
	    static createFrom(source: any = {}) {
	        return new Keycode(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.Rawcode = source["Rawcode"];
	        this.Display = source["Display"];
	        this.JSCode = source["JSCode"];
	    }
	}

}

export namespace soundhotkey {
	
	export class SoundHotkey {
	    id: number[];
	    path: string;
	    hotkey: number[];
	
	    static createFrom(source: any = {}) {
	        return new SoundHotkey(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.id = source["id"];
	        this.path = source["path"];
	        this.hotkey = source["hotkey"];
	    }
	}

}

