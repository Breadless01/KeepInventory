export namespace backend {
	
	export class CreateBauteilRequest {
	    TeilName: string;
	    KundeID: number;
	    ProjektID: number;
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
	        this.KundeID = source["KundeID"];
	        this.ProjektID = source["ProjektID"];
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
	    kunden_id: number;
	
	    static createFrom(source: any = {}) {
	        return new CreateProjektRequest(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.name = source["name"];
	        this.kunden_id = source["kunden_id"];
	    }
	}

}

export namespace domain {
	
	export class Bauteil {
	    ID: number;
	    TeilName: string;
	    KundeID: number;
	    ProjektID: number;
	    // Go type: time
	    Erstelldatum: any;
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
	        this.KundeID = source["KundeID"];
	        this.ProjektID = source["ProjektID"];
	        this.Erstelldatum = this.convertValues(source["Erstelldatum"], null);
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
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice && a.map) {
		        return (a as any[]).map(elem => this.convertValues(elem, classs));
		    } else if ("object" === typeof a) {
		        if (asMap) {
		            for (const key of Object.keys(a)) {
		                a[key] = new classs(a[key]);
		            }
		            return a;
		        }
		        return new classs(a);
		    }
		    return a;
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
	export class Projekt {
	    ID: number;
	    Name: string;
	    KundeID: number;
	    Kunde?: Kunde;
	
	    static createFrom(source: any = {}) {
	        return new Projekt(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.ID = source["ID"];
	        this.Name = source["Name"];
	        this.KundeID = source["KundeID"];
	        this.Kunde = this.convertValues(source["Kunde"], Kunde);
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice && a.map) {
		        return (a as any[]).map(elem => this.convertValues(elem, classs));
		    } else if ("object" === typeof a) {
		        if (asMap) {
		            for (const key of Object.keys(a)) {
		                a[key] = new classs(a[key]);
		            }
		            return a;
		        }
		        return new classs(a);
		    }
		    return a;
		}
	}

}

