{ config, lib, pkgs, ... }:
let
  cfg = config.services.zing;
in
{
  options.services.zing = {
    enable = lib.mkEnableOption "Enable Zing service";

    package = lib.mkOption {
      type = lib.types.package;
      default = pkgs.zing;
      description = "Package providing the zing binary";
    };

    user = lib.mkOption { type = lib.types.str; default = "zing"; description = "User running the service"; };
    group = lib.mkOption { type = lib.types.str; default = "zing"; description = "Group running the service"; };
    dataDir = lib.mkOption { type = lib.types.path; default = "/var/lib/zing"; description = "State directory"; };
    port = lib.mkOption { type = lib.types.port; default = 8080; description = "Port to listen on"; };
    extraArgs = lib.mkOption { type = lib.types.listOf lib.types.str; default = [ ]; description = "Extra CLI args"; };
    environment = lib.mkOption { type = lib.types.attrsOf lib.types.str; default = { }; description = "Environment variables"; };
  };

  config = lib.mkIf cfg.enable {
    users.users.${cfg.user} = {
      isSystemUser = true;
      group = cfg.group;
      home = cfg.dataDir;
    };
    users.groups.${cfg.group} = { };

    systemd.services.zing = {
      description = "Zing Service";
      wantedBy = [ "multi-user.target" ];
      after = [ "network-online.target" ];
      wants = [ "network-online.target" ];

      serviceConfig = {
        ExecStart = ''${cfg.package}/bin/zing --port ${toString cfg.port} ${lib.concatStringsSep " " cfg.extraArgs}'';
        User = cfg.user;
        Group = cfg.group;
        Restart = "on-failure";
        WorkingDirectory = cfg.dataDir;
        StateDirectory = [ (baseNameOf cfg.dataDir) ];
        StateDirectoryMode = "0750";
        # Hardening (tune as needed)
        AmbientCapabilities = "";
        CapabilityBoundingSet = "";
        DevicePolicy = "closed";
        LockPersonality = true;
        MemoryDenyWriteExecute = true;
        NoNewPrivileges = true;
        PrivateDevices = true;
        PrivateTmp = true;
        ProtectClock = true;
        ProtectControlGroups = true;
        ProtectHome = true;
        ProtectHostname = true;
        ProtectKernelLogs = true;
        ProtectKernelModules = true;
        ProtectKernelTunables = true;
        ProtectSystem = "strict";
        RestrictSUIDSGID = true;
        RestrictAddressFamilies = [ "AF_INET" "AF_INET6" "AF_UNIX" ];
        SystemCallArchitectures = "native";
        UMask = "0027";
      };

      environment = cfg.environment;
    };
  };
}
