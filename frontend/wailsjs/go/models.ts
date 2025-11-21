export namespace backend {
	
	export class CreateBauteilRequest {
	    TeilName: string;
	    KundeId: number;
	    ProjektId: number;
	    TypID: number;
	    HerstellungsartID: number;
	    VerschleissteilID: number;
	    FunktionID: number;
	    MaterialID: number;
	    OberflaechenbehandlungID: number;
	    FarbeID: number;
	    ReserveID: number;
	
	    static createFrom(source: any = {}) {
	        return new CreateBauteilRequest(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.TeilName = source["TeilName"];
	        this.KundeId = source["KundeId"];
	        this.ProjektId = source["ProjektId"];
	        this.TypID = source["TypID"];
	        this.HerstellungsartID = source["HerstellungsartID"];
	        this.VerschleissteilID = source["VerschleissteilID"];
	        this.FunktionID = source["FunktionID"];
	        this.MaterialID = source["MaterialID"];
	        this.OberflaechenbehandlungID = source["OberflaechenbehandlungID"];
	        this.FarbeID = source["FarbeID"];
	        this.ReserveID = source["ReserveID"];
	    }
	}
	export class CreateKundeRequest {
	    name: string;
	    sitz: string;
	
	    static createFrom(source: any = {}) {
	        return new CreateKundeRequest(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.name = source["name"];
	        this.sitz = source["sitz"];
	    }
	}
	export class CreateProjektRequest {
	    name: string;
	    kunde: string;
	
	    static createFrom(source: any = {}) {
	        return new CreateProjektRequest(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.name = source["name"];
	        this.kunde = source["kunde"];
	    }
	}

}

export namespace domain {
	
	export class Bauteil {
	    ID: number;
	    TeilName: string;
	    Kunde: string;
	    KundeId: number;
	    Projekt: string;
	    ProjektId: number;
	    Erstelldatum: string;
	    TypID: number;
	    HerstellungsartID: number;
	    VerschleissteilID: number;
	    FunktionID: number;
	    MaterialID: number;
	    OberflaechenbehandlungID: number;
	    FarbeID: number;
	    ReserveID: number;
	    Sachnummer: string;
	
	    static createFrom(source: any = {}) {
	        return new Bauteil(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.ID = source["ID"];
	        this.TeilName = source["TeilName"];
	        this.Kunde = source["Kunde"];
	        this.KundeId = source["KundeId"];
	        this.Projekt = source["Projekt"];
	        this.ProjektId = source["ProjektId"];
	        this.Erstelldatum = source["Erstelldatum"];
	        this.TypID = source["TypID"];
	        this.HerstellungsartID = source["HerstellungsartID"];
	        this.VerschleissteilID = source["VerschleissteilID"];
	        this.FunktionID = source["FunktionID"];
	        this.MaterialID = source["MaterialID"];
	        this.OberflaechenbehandlungID = source["OberflaechenbehandlungID"];
	        this.FarbeID = source["FarbeID"];
	        this.ReserveID = source["ReserveID"];
	        this.Sachnummer = source["Sachnummer"];
	    }
	}
	export class Farbe {
	    ID: number;
	    Name: string;
	    Symbol: number;
	
	    static createFrom(source: any = {}) {
	        return new Farbe(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.ID = source["ID"];
	        this.Name = source["Name"];
	        this.Symbol = source["Symbol"];
	    }
	}
	export class Funktion {
	    ID: number;
	    Name: string;
	    Symbol: number;
	
	    static createFrom(source: any = {}) {
	        return new Funktion(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.ID = source["ID"];
	        this.Name = source["Name"];
	        this.Symbol = source["Symbol"];
	    }
	}
	export class Herstellungsart {
	    ID: number;
	    Name: string;
	    Symbol: number;
	
	    static createFrom(source: any = {}) {
	        return new Herstellungsart(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.ID = source["ID"];
	        this.Name = source["Name"];
	        this.Symbol = source["Symbol"];
	    }
	}
	export class Kunde {
	    ID: number;
	    Name: string;
	    Sitz: string;
	
	    static createFrom(source: any = {}) {
	        return new Kunde(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.ID = source["ID"];
	        this.Name = source["Name"];
	        this.Sitz = source["Sitz"];
	    }
	}
	export class Material {
	    ID: number;
	    Name: string;
	    Symbol: number;
	
	    static createFrom(source: any = {}) {
	        return new Material(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.ID = source["ID"];
	        this.Name = source["Name"];
	        this.Symbol = source["Symbol"];
	    }
	}
	export class Oberflaechenbehandlung {
	    ID: number;
	    Name: string;
	    Symbol: number;
	
	    static createFrom(source: any = {}) {
	        return new Oberflaechenbehandlung(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.ID = source["ID"];
	        this.Name = source["Name"];
	        this.Symbol = source["Symbol"];
	    }
	}
	export class Projekt {
	    ID: number;
	    Name: string;
	    Kunde: string;
	
	    static createFrom(source: any = {}) {
	        return new Projekt(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.ID = source["ID"];
	        this.Name = source["Name"];
	        this.Kunde = source["Kunde"];
	    }
	}
	export class Reserve {
	    ID: number;
	    Name: string;
	    Symbol: number;
	
	    static createFrom(source: any = {}) {
	        return new Reserve(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.ID = source["ID"];
	        this.Name = source["Name"];
	        this.Symbol = source["Symbol"];
	    }
	}
	export class Typ {
	    ID: number;
	    Name: string;
	    Symbol: number;
	
	    static createFrom(source: any = {}) {
	        return new Typ(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.ID = source["ID"];
	        this.Name = source["Name"];
	        this.Symbol = source["Symbol"];
	    }
	}
	export class Verschleissteil {
	    ID: number;
	    Name: string;
	    Symbol: number;
	
	    static createFrom(source: any = {}) {
	        return new Verschleissteil(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.ID = source["ID"];
	        this.Name = source["Name"];
	        this.Symbol = source["Symbol"];
	    }
	}

}

