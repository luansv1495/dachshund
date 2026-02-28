import { LucideIcon, PlayIcon, PlusIcon } from "lucide-react";
import { useState } from "react";
import { Button } from "./ui/button";

export function TabBar() {
  const [tabs, setTabs] = useState<
    {
      title: string;
      type: string;
      metadata?: string;
    }[]
  >([
    { title: "Query 1", type: "query" },
    { title: "Table 1", type: "table" },
    { title: "Table 1 - Struct", type: "table-struct" },
    { title: "Table 1 - Diagram", type: "table-diagram" },
    { title: "Schema 1 - Diagram", type: "schema-diagram" },
    { title: "Schema 1 - Struct", type: "schema-struct" },
    { title: "Database 1 - Dashboard", type: "database-dashboard" },
  ]);

  return (
    <header className="bg-background sticky top-0 flex h-12 shrink-0 items-center gap-2 border-b px-4">
      {tabs.map((tab, i) => {
        return <div key={i}>{tab.title}</div>;
      })}
      <Button variant="outline" size="icon">
        <PlusIcon />
      </Button>
    </header>
  );
}
