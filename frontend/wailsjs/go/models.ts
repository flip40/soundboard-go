export namespace audiodevice {
	
	export class AudioDevice {
	    ID: string;
	    Name: string;
	
	    static createFrom(source: any = {}) {
	        return new AudioDevice(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.ID = source["ID"];
	        this.Name = source["Name"];
	    }
	}

}

export namespace soundhotkey {
	
	export class SoundHotkey {
	    ID: number[];
	    Path: string;
	    Hotkey: number[];
	
	    static createFrom(source: any = {}) {
	        return new SoundHotkey(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.ID = source["ID"];
	        this.Path = source["Path"];
	        this.Hotkey = source["Hotkey"];
	    }
	}

}

