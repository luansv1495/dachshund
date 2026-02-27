import {
  Dialog,
  DialogContent,
  DialogDescription,
  DialogHeader,
  DialogTitle,
  DialogTrigger,
} from "@/components/ui/dialog";
import { Button } from "./ui/button";
import { PlusIcon } from "lucide-react";
import { Field, FieldGroup, FieldLabel } from "./ui/field";
import {
  Select,
  SelectContent,
  SelectItem,
  SelectTrigger,
  SelectValue,
} from "./ui/select";
import { useState } from "react";
import { core } from "../../wailsjs/go/models";
import { Input } from "./ui/input";
import { CreateConnection, TestConnection } from "../../wailsjs/go/main/App";
import { useMutation, useQueryClient } from "@tanstack/react-query";

export function CreateConnectionForm({
  connectionTypes,
}: {
  connectionTypes: core.ConnectionType[];
}) {
  const queryClient = useQueryClient();
  const [open, setOpen] = useState(false);
  const [connectionData, setConnectionData] = useState<{
    type: core.ConnectionType;
    host: string;
    port: number;
    user?: string;
    password?: string;
    database?: string;
  }>({
    type: connectionTypes[0],
    host: "localhost",
    port: connectionTypes[0].defaultPort,
    user: "",
    password: "",
    database: "",
  });

  const handleTest = async () => {
    try {
      await TestConnection({
        ID: "test",
        Type: connectionData.type.id,
        Host: connectionData.host,
        Port: connectionData.port,
        User: connectionData.user ?? "",
        Password: connectionData.password ?? "",
        Database: connectionData.database ?? "",
        SSLMode: "",
      });
      alert("Connection successful!");
    } catch (err) {
      alert("Connection failed: " + err);
    }
  };

  const handleConnect = async () => {
    const conn = await CreateConnection({
      ID: "",
      Type: connectionData.type.id,
      Host: connectionData.host,
      Port: connectionData.port,
      User: connectionData.user ?? "",
      Password: connectionData.password ?? "",
      Database: connectionData.database ?? "",
      SSLMode: "",
    });
    localStorage.setItem("connection", conn.ID);
    setOpen(false);
  };

  const mutation = useMutation({
    mutationFn: handleConnect,
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: ["connections"] });
      queryClient.invalidateQueries({ queryKey: ["connection"] });
      queryClient.invalidateQueries({ queryKey: ["tree"] });
    },
  });

  return (
    <Dialog onOpenChange={setOpen} open={open}>
      <DialogTrigger asChild>
        <Button variant="ghost">
          <div className="flex size-6 items-center justify-center rounded-md border bg-transparent">
            <PlusIcon className="size-4" />
          </div>{" "}
          Create Connection
        </Button>
      </DialogTrigger>
      <DialogContent>
        <DialogHeader>
          <DialogTitle>New Connection</DialogTitle>
          <DialogDescription>
            This action cannot be undone. This will permanently delete your
            account and remove your data from our servers.
          </DialogDescription>
        </DialogHeader>
        <form>
          <FieldGroup>
            <Field>
              <FieldLabel>Connection Type</FieldLabel>
              <Select
                value={connectionData.type.id}
                onValueChange={(id) => {
                  const type = connectionTypes.find((t) => t.id == id)!;

                  setConnectionData({
                    ...connectionData,
                    type: type,
                    port: type.defaultPort,
                  });
                }}
              >
                <SelectTrigger>
                  <SelectValue placeholder="Choose connection type" />
                </SelectTrigger>
                <SelectContent>
                  {connectionTypes.map((type) => {
                    return (
                      <SelectItem key={type.id} value={type.id}>
                        {type.name}
                      </SelectItem>
                    );
                  })}
                </SelectContent>
              </Select>
            </Field>
            <div className="grid grid-cols-4 gap-4">
              <Field className="col-span-3">
                <FieldLabel htmlFor="host">Host</FieldLabel>
                <Input
                  id="host"
                  type="text"
                  defaultValue={connectionData.host}
                  onChange={(e) =>
                    setConnectionData({
                      ...connectionData,
                      host: e.target.value,
                    })
                  }
                />
              </Field>
              <Field>
                <FieldLabel htmlFor="port">Port</FieldLabel>
                <Input
                  id="port"
                  type="number"
                  value={connectionData.port}
                  onChange={(e) =>
                    setConnectionData({
                      ...connectionData,
                      port: Number(e.target.value),
                    })
                  }
                />
              </Field>
            </div>
            <div className="grid grid-cols-2 gap-4">
              <Field>
                <FieldLabel htmlFor="user">User</FieldLabel>
                <Input
                  id="user"
                  type="text"
                  defaultValue={connectionData.user}
                  onChange={(e) =>
                    setConnectionData({
                      ...connectionData,
                      user: e.target.value,
                    })
                  }
                />
              </Field>
              <Field>
                <FieldLabel htmlFor="password">Password</FieldLabel>
                <Input
                  id="password"
                  type="password"
                  value={connectionData.password}
                  onChange={(e) =>
                    setConnectionData({
                      ...connectionData,
                      password: e.target.value,
                    })
                  }
                />
              </Field>
            </div>
            <Field className="col-span-3">
              <FieldLabel htmlFor="database">Database</FieldLabel>
              <Input
                id="database"
                type="text"
                defaultValue={connectionData.database}
                onChange={(e) =>
                  setConnectionData({
                    ...connectionData,
                    database: e.target.value,
                  })
                }
              />
            </Field>
            <Field orientation="horizontal">
              <Button type="button" onClick={() => mutation.mutate()}>
                Connect
              </Button>
              <Button variant="outline" type="button" onClick={handleTest}>
                Test
              </Button>
            </Field>
          </FieldGroup>
        </form>
      </DialogContent>
    </Dialog>
  );
}
