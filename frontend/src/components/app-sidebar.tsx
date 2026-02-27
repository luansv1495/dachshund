import * as React from "react";

import {
  Sidebar,
  SidebarContent,
  SidebarHeader,
  SidebarRail,
} from "@/components/ui/sidebar";
import { ConnectionSwitcher } from "./connection-switcher";
import { core } from "../../wailsjs/go/models";
import { useQuery } from "@tanstack/react-query";
import { TreeBuilder } from "./tree";

export function AppSidebar({
  connections,
  connectionTypes,
  ...props
}: {
  connections: core.ConnectionConfig[];
  connectionTypes: core.ConnectionType[];
} & React.ComponentProps<typeof Sidebar>) {
  const { data: connId } = useQuery({
    queryKey: ["connection"],
    queryFn: () => localStorage.getItem("connection"),
  });

  return (
    <Sidebar {...props}>
      <SidebarHeader>
        <ConnectionSwitcher
          connections={connections}
          connectionTypes={connectionTypes}
          defaultConnection={connections.find((conn) => conn.ID == connId)}
        />
      </SidebarHeader>
      <SidebarContent className="gap-0">
        {!!connId && <TreeBuilder connId={connId} />}
      </SidebarContent>
      <SidebarRail />
    </Sidebar>
  );
}
