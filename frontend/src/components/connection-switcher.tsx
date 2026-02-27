"use client";

import * as React from "react";
import {
  Check,
  ChevronsUpDown,
  DatabaseIcon,
  GalleryVerticalEnd,
  PlusIcon,
} from "lucide-react";

import {
  DropdownMenu,
  DropdownMenuContent,
  DropdownMenuItem,
  DropdownMenuSeparator,
  DropdownMenuTrigger,
} from "@/components/ui/dropdown-menu";
import {
  SidebarMenu,
  SidebarMenuButton,
  SidebarMenuItem,
} from "@/components/ui/sidebar";
import { CreateConnectionForm } from "./create-connection-form";

import { core } from "../../wailsjs/go/models";

export function ConnectionSwitcher({
  connections,
  connectionTypes,
  defaultConnection,
}: {
  connections: core.ConnectionConfig[];
  connectionTypes: core.ConnectionType[];
  defaultConnection?: core.ConnectionConfig;
}) {
  const [selectedConn, setSelectedConn] = React.useState(defaultConnection);

  React.useEffect(() => {
    if (!!defaultConnection) {
      setSelectedConn(defaultConnection);
    }
  }, [defaultConnection]);

  if (connections.length == 0 && connectionTypes.length > 0) {
    return (
      <SidebarMenu>
        <SidebarMenuItem>
          <CreateConnectionForm connectionTypes={connectionTypes} />
        </SidebarMenuItem>
      </SidebarMenu>
    );
  }

  return (
    <SidebarMenu>
      <SidebarMenuItem>
        <DropdownMenu>
          <DropdownMenuTrigger asChild>
            <SidebarMenuButton
              size="lg"
              className="data-[state=open]:bg-sidebar-accent data-[state=open]:text-sidebar-accent-foreground"
            >
              <div className="bg-sidebar-primary text-sidebar-primary-foreground flex aspect-square size-8 items-center justify-center rounded-lg">
                <DatabaseIcon className="size-4" />
              </div>
              <div className="flex flex-col gap-0.5 leading-none">
                <span className="font-medium">{selectedConn?.Type}</span>
                <span>{selectedConn?.Database}</span>
              </div>
              <ChevronsUpDown className="ml-auto" />
            </SidebarMenuButton>
          </DropdownMenuTrigger>
          <DropdownMenuContent
            className="w-(--radix-dropdown-menu-trigger-width)"
            align="start"
          >
            {connections.map((conn, i) => (
              <DropdownMenuItem
                key={i}
                onSelect={() => {
                  localStorage.setItem("connection", conn.ID);
                  setSelectedConn(conn);
                }}
              >
                {conn.Database}{" "}
                {conn.ID === selectedConn?.ID && <Check className="ml-auto" />}
              </DropdownMenuItem>
            ))}
            <DropdownMenuSeparator />
            <DropdownMenuItem onSelect={(e) => e.preventDefault()}>
              <CreateConnectionForm connectionTypes={connectionTypes} />
            </DropdownMenuItem>
          </DropdownMenuContent>
        </DropdownMenu>
      </SidebarMenuItem>
    </SidebarMenu>
  );
}
