import { AppSidebar } from "./components/app-sidebar";
import { SidebarInset, SidebarProvider } from "./components/ui/sidebar";
import { ListConnections, ListConnectionTypes } from "../wailsjs/go/main/App";
import { core } from "../wailsjs/go/models";

import "./App.css";
import { useQuery } from "@tanstack/react-query";

function App() {
  const { data: connections } = useQuery({
    queryKey: ["connections"],
    queryFn: ListConnections,
  });
  const { data: connectionTypes } = useQuery({
    queryKey: ["connection_types"],
    queryFn: ListConnectionTypes,
  });

  return (
    <SidebarProvider
      style={
        {
          "--sidebar-width": "20rem",
          "--sidebar-width-mobile": "20rem",
        } as React.CSSProperties
      }
    >
      <AppSidebar
        connections={connections ?? []}
        connectionTypes={connectionTypes ?? []}
      />
      <SidebarInset>
        <header className="bg-background sticky top-0 flex h-10 shrink-0 items-center gap-2 border-b px-4"></header>
        <div className="flex flex-1 flex-col gap-4 p-4">
          {Array.from({ length: 24 }).map((_, index) => (
            <div
              key={index}
              className="bg-muted/50 aspect-video h-12 w-full rounded-lg"
            />
          ))}
        </div>
      </SidebarInset>
    </SidebarProvider>
  );
}

export default App;
