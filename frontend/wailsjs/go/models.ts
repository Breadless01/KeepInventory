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
	export class BauteilFilterResult {
	    items: Bauteil[];
	    total: number;
	    facets: Record<string, Array<FacetOption>>;
	
	    static createFrom(source: any = {}) {
	        return new BauteilFilterResult(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.items = this.convertValues(source["items"], Bauteil);
	        this.total = source["total"];
	        this.facets = this.convertValues(source["facets"], Array<FacetOption>, true);
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
	export class FacetOption {
	    id: number;
	    name: string;
	    count: number;
	
	    static createFrom(source: any = {}) {
	        return new FacetOption(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.id = source["id"];
	        this.name = source["name"];
	        this.count = source["count"];
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
	export class FieldFilterConfig {
	    field: string;
	    label: string;
	    enabled: boolean;
	
	    static createFrom(source: any = {}) {
	        return new FieldFilterConfig(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.field = source["field"];
	        this.label = source["label"];
	        this.enabled = source["enabled"];
	    }
	}
	export class ResourceFilterConfig {
	    resource: string;
	    table: string;
	    fields: FieldFilterConfig[];
	
	    static createFrom(source: any = {}) {
	        return new ResourceFilterConfig(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.resource = source["resource"];
	        this.table = source["table"];
	        this.fields = this.convertValues(source["fields"], FieldFilterConfig);
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
	export class FilterConfig {
	    resources: ResourceFilterConfig[];
	
	    static createFrom(source: any = {}) {
	        return new FilterConfig(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.resources = this.convertValues(source["resources"], ResourceFilterConfig);
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
	export class FilterState {
	    resource: string;
	    page: number;
	    pageSize: number;
	    filters: Record<string, Array<any>>;
	
	    static createFrom(source: any = {}) {
	        return new FilterState(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.resource = source["resource"];
	        this.page = source["page"];
	        this.pageSize = source["pageSize"];
	        this.filters = source["filters"];
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
	export class KundeFilterResult {
	    items: Kunde[];
	    total: number;
	    facets: Record<string, Array<FacetOption>>;
	
	    static createFrom(source: any = {}) {
	        return new KundeFilterResult(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.items = this.convertValues(source["items"], Kunde);
	        this.total = source["total"];
	        this.facets = this.convertValues(source["facets"], Array<FacetOption>, true);
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
	export class ProjektFilterResult {
	    items: Projekt[];
	    total: number;
	    facets: Record<string, Array<FacetOption>>;
	
	    static createFrom(source: any = {}) {
	        return new ProjektFilterResult(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.items = this.convertValues(source["items"], Projekt);
	        this.total = source["total"];
	        this.facets = this.convertValues(source["facets"], Array<FacetOption>, true);
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
	
	export class SearchResult {
	    id: number;
	    type: string;
	    label: string;
	    subtitle: string;
	    extra: Record<string, string>;
	
	    static createFrom(source: any = {}) {
	        return new SearchResult(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.id = source["id"];
	        this.type = source["type"];
	        this.label = source["label"];
	        this.subtitle = source["subtitle"];
	        this.extra = source["extra"];
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

