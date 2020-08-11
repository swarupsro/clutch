module.exports = {
  "@clutch-sh/ec2": {
    terminateInstance: {
      trending: true,
      componentProps: {
        resolverType: "clutch.aws.ec2.v1.Instance",
      },
    },
    resizeAutoscalingGroup: {
      componentProps: {
        resolverType: "clutch.aws.ec2.v1.AutoscalingGroup",
        notes: [
          {
            severity: "info",
            text:
              "The autoscaling group may take several minutes to bring additional instances online.",
          },
        ],
      },
    },
  },
  "@clutch-sh/envoy": {
    remoteTriage: {
      trending: true,
      componentProps: {
        options: {
          Clusters: "clusters",
          Listeners: "listeners",
          Runtime: "runtime",
          Stats: "stats",
          "Server Info": "serverInfo",
        },
      },
    },
  },
  "@clutch-sh/k8s": {
    deletePod: {
      componentProps: {
        resolverType: "clutch.k8s.v1.Pod",
      },
    },
    resizeHPA: {
      trending: true,
      componentProps: {
        resolverType: "clutch.k8s.v1.HPA",
      },
    },
  },
  "@clutch-sh/experimentation": {
    listExperiments: {
      trending: true,
      componentProps: {
        columns: ["identifier", "targets", "type", "description"],
        experimentTypes: {
          "type.googleapis.com/clutch.chaos.serverexperimentation.v1.TestSpecification": {
            mapping: function(e) {
              const config = e.latency || e.abort;
              const clusters = config.clusterPair
              return {
                "targets": `from "${clusters.downstreamCluster}" to "${clusters.upstreamCluster}"`,
                "type": "server",
                "description": e.abort ? `${e.abort.percent}% abort with ${e.abort.httpStatus} status code` : `${e.latency.percent}% latency with ${e.latency.durationMs}ms delay`,
              }
            },
            links: [
              {
                text: "Start Abort",
                path: "/serverexperimentation/startabort",
              },
              {
                text: "Start Latency",
                path: "/serverexperimentation/startlatency",
              },
            ]
          }
        }
      },
    },
  },
  "@clutch-sh/serverexperimentation": {
    startAbortExperiment: {},
    startLatencyExperiment: {},
  },
};
