export namespace core {
	
	export class ConnectionConfig {
	    ID: string;
	    Type: string;
	    Host: string;
	    Port: number;
	    User: string;
	    Password: string;
	    Database: string;
	    SSLMode: string;
	
	    static createFrom(source: any = {}) {
	        return new ConnectionConfig(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.ID = source["ID"];
	        this.Type = source["Type"];
	        this.Host = source["Host"];
	        this.Port = source["Port"];
	        this.User = source["User"];
	        this.Password = source["Password"];
	        this.Database = source["Database"];
	        this.SSLMode = source["SSLMode"];
	    }
	}
	export class ConnectionType {
	    id: string;
	    name: string;
	    defaultPort: number;
	
	    static createFrom(source: any = {}) {
	        return new ConnectionType(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.id = source["id"];
	        this.name = source["name"];
	        this.defaultPort = source["defaultPort"];
	    }
	}
	export class TreeNode {
	    id: string;
	    name: string;
	    type: string;
	    metadata: string;
	    hasChildren: boolean;
	
	    static createFrom(source: any = {}) {
	        return new TreeNode(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.id = source["id"];
	        this.name = source["name"];
	        this.type = source["type"];
	        this.metadata = source["metadata"];
	        this.hasChildren = source["hasChildren"];
	    }
	}

}

