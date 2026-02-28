import { AppSidebar } from "./components/app-sidebar";
import { SidebarInset, SidebarProvider } from "./components/ui/sidebar";
import { ListConnections, ListConnectionTypes } from "../wailsjs/go/main/App";
import { core } from "../wailsjs/go/models";

import "./App.css";
import { useQuery } from "@tanstack/react-query";
import { TabBar } from "./components/tab-bar";

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
          "--sidebar-width": "25rem",
          "--sidebar-width-mobile": "25rem",
        } as React.CSSProperties
      }
    >
      <AppSidebar
        connections={connections ?? []}
        connectionTypes={connectionTypes ?? []}
      />
      <SidebarInset>
        <TabBar />
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
